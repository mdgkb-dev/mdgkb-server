package models

type ChatRoom struct {
	chatID        string
	clients       map[*ChatMember]bool
	broadcastChan chan ChatMessage
	registerChan  chan *ChatMember
	exitChan      chan *ChatMember
}

type ChatRooms []*ChatRoom

var chatRoomsPool = NewChatRooms()

func NewChatRooms() ChatRooms {
	return make(ChatRooms, 0)
}

func GetChatRoom(chatID string) *ChatRoom {
	var foundChat *ChatRoom
	for i := range chatRoomsPool {
		if chatRoomsPool[i].chatID == chatID {
			foundChat = chatRoomsPool[i]
			break
		}
	}
	if foundChat == nil {
		foundChat = newChatRoom(chatID)
		go foundChat.run()
		chatRoomsPool = append(chatRoomsPool, foundChat)
	}
	return foundChat
}

func newChatRoom(chatID string) *ChatRoom {
	return &ChatRoom{
		chatID:        chatID,
		broadcastChan: make(chan ChatMessage),
		registerChan:  make(chan *ChatMember),
		exitChan:      make(chan *ChatMember),
		clients:       make(map[*ChatMember]bool),
	}
}

func (h *ChatRoom) run() {
	for {
		select {
		case client := <-h.registerChan:
			m := ChatMessage{Message: "К чату подключились", Type: enter}
			for client := range h.clients {
				client.send <- m
			}
			h.clients[client] = true

		case client := <-h.exitChan:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
			m := ChatMessage{Message: "Из чата вышли", Type: exit}
			for client := range h.clients {
				client.send <- m
			}

		case chatMessage := <-h.broadcastChan:

			for client := range h.clients {
				select {
				case client.send <- chatMessage:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
