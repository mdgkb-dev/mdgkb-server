package broker

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// the amount of time to wait when pushing a message to
// a slow client or a client that closed after `range clients` started.
const patience time.Duration = time.Second * 1

type (
	NotificationEvent struct {
		EventName string
		Payload   interface{}
	}

	NotifierChan chan NotificationEvent

	Broker struct {

		// Events are pushed to this channel by the main events-gathering routine
		Notifier NotifierChan

		// New client connections
		newClients chan NotifierChan

		// Closed client connections
		closingClients chan NotifierChan

		// Client connections registry
		clients map[NotifierChan]struct{}
	}
)

func NewBroker() (broker *Broker) {
	// Instantiate a broker
	b := &Broker{
		Notifier:       make(NotifierChan, 1),
		newClients:     make(chan NotifierChan),
		closingClients: make(chan NotifierChan),
		clients:        make(map[NotifierChan]struct{}),
	}
	//b.Listen()
	return b
}

func (broker *Broker) ServeHTTP(c *gin.Context) {
	eventName := c.Param("channel")
	fmt.Println(eventName)
	fmt.Println(eventName)
	fmt.Println(eventName)
	fmt.Println(eventName)
	fmt.Println(eventName)
	fmt.Println(eventName)
	fmt.Println(eventName)
	fmt.Println(eventName)
	fmt.Println(eventName)
	fmt.Println(eventName)
	fmt.Println(eventName)
	//c.Header("Content-Type", "text/event-stream")
	//c.Header("Cache-Control", "no-cache")
	//c.Header("Connection", "keep-alive")
	//c.Header("Access-Control-Allow-Origin", "*")

	messageChan := make(NotifierChan)
	broker.newClients <- messageChan
	fmt.Println(broker.newClients)
	fmt.Println(broker.newClients)
	fmt.Println(broker.newClients)
	fmt.Println(broker.newClients)
	defer func() {
		broker.closingClients <- messageChan
	}()

	c.Stream(func(w io.Writer) bool {
		event := <-messageChan
		fmt.Println("123123sdfsdf")
		fmt.Println("123123sdfsdf")
		fmt.Println("123123sdfsdf")
		fmt.Println("123123sdfsdf")
		fmt.Println("123123sdfsdf")
		switch eventName {
		case event.EventName:
			c.SSEvent(event.EventName, event.Payload)
		}
		c.Writer.Flush()
		return true
	})
}

// Listen for new notifications and redistribute them to clients
func (broker *Broker) Listen() {
	for {
		select {
		case s := <-broker.newClients:
			broker.clients[s] = struct{}{}
		case s := <-broker.closingClients:
			delete(broker.clients, s)
		case event := <-broker.Notifier:
			for clientMessageChan := range broker.clients {
				select {
				case clientMessageChan <- event:
				case <-time.After(patience):
					log.Print("Skipping client.")
				}
			}
		}
	}
}
