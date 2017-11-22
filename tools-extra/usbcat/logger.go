package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

type Message struct {
	Type string
	Text string
}

type Logger interface {
	Connected()
	Disconnected()
	Log(msg Message)
}

// Dumb mode (STDOUT logger)
type StdoutLogger struct{}

func (s StdoutLogger) Connected() {
	fmt.Fprintln(os.Stderr, "==============================")
}

func (s StdoutLogger) Disconnected() {
	fmt.Fprintln(os.Stderr, "==============================")
}

func (s StdoutLogger) Log(msg Message) {
	fmt.Printf("[%s] %s\n", msg.Type, msg.Text)
}

var upgrader = websocket.Upgrader{}

// Metrics mode (Websocket logger)
type WebsocketLogger struct {
	connected bool
	messages  chan Message
	pump      *DynamicFanOut
}

func MakeWSLogger() *WebsocketLogger {
	messages := make(chan Message, 16)
	return &WebsocketLogger{
		connected: false,
		messages:  messages,
		pump:      runDFO(messages),
	}
}

func (l *WebsocketLogger) Connected() {
	l.connected = true
	l.messages <- Message{
		Type: "connection",
		Text: "connected",
	}
}

func (l *WebsocketLogger) Disconnected() {
	l.connected = false
	l.messages <- Message{
		Type: "connection",
		Text: "disconnected",
	}
}

func (l *WebsocketLogger) Log(msg Message) {
	l.messages <- msg
}

func (l *WebsocketLogger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WS:", err)
		return
	}
	defer c.Close()

	// Subscribe to pump
	messages := l.pump.Subscribe()
	defer l.pump.Unsubscribe(messages)

	for msg := range messages {
		err = c.WriteJSON(msg)
		if err != nil {
			log.Println("Error writing to client:", err)
			break
		}
	}
}
