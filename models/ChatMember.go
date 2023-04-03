package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type ChatMember struct {
	userID   string
	chatRoom *ChatRoom
	conn     *websocket.Conn
	send     chan ChatMessage
}

func NewChatMember(conn *websocket.Conn, userID string) *ChatMember {
	return &ChatMember{conn: conn, send: make(chan ChatMessage, 256), userID: userID}
}

func (item *ChatMember) StartChat(chatID string) {
	item.enterToChat(chatID)

	go item.writePump()
	go item.readPump()

	item.chatRoom = GetChatRoom(chatID)
	item.chatRoom.registerChan <- item
	//item.send <- []byte("Welcome")
}

func (item *ChatMember) enterToChat(chatID string) {
	item.chatRoom = GetChatRoom(chatID)
}

const (
	writeWait      = 600 * time.Second
	pongWait       = 600 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512000
)

func (item *ChatMember) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		item.conn.Close()
	}()
	for {
		select {
		case message, ok := <-item.send:
			_ = item.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				_ = item.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := item.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			_, err = w.Write(message.ToBytes())
			if err != nil {
				return
			}
			// Add queued chat messages to the current websocket message.
			n := len(item.send)
			for i := 0; i < n; i++ {
				_, _ = w.Write(newline)
				b := <-item.send
				_, _ = w.Write(b.ToBytes())
			}
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			_ = item.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := item.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

var (
	newline = []byte{'\n'}
	//space   = []byte{' '}
)

func (item *ChatMember) readPump() {
	defer func() {
		item.chatRoom.exitChan <- item
		_ = item.conn.Close()
	}()
	item.conn.SetReadLimit(maxMessageSize)
	_ = item.conn.SetReadDeadline(time.Now().Add(pongWait))
	item.conn.SetPongHandler(func(string) error {
		_ = item.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	for {
		_, message, err := item.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		//m := ChatMessage{}
		//_ = json.Unmarshal(bytes.TrimSpace(bytes.Replace(message, newline, space, -1)), &m)
		//data := map[string][]byte{
		//	"chatId": m.ChatID,
		//	"chatId": m.ChatID,
		//}
		//userMessage, _ := json.Marshal(m)
		var chatMessage ChatMessage
		_ = json.Unmarshal(message, &chatMessage)
		if chatMessage.IsPing() {
			continue
		}
		item.chatRoom.broadcastChan <- chatMessage
	}
}
