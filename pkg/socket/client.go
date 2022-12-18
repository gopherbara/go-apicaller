package socket

import (
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

type Client struct {
	mu      sync.Mutex
	Id      string
	Manager *Manager
	Conn    *websocket.Conn
}

func NewClient(manager *Manager, con *websocket.Conn) *Client {
	return &Client{
		Manager: manager,
		Conn:    con,
	}
}

func (cl *Client) Read() {
	// unregister client in the end
	defer func() {
		cl.Manager.Unregister <- cl
		cl.Conn.Close()
	}()

	for {
		_, _, err := cl.Conn.ReadMessage()
		if err != nil {
			// not fatal cause server stop - not sure why
			log.Printf("error on read from client: %s", err.Error())
			return
		}
	}
}
