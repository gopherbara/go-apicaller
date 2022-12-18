package socket

import (
	"log"
	"sync"
)

type Manager struct {
	Register      chan *Client
	Unregister    chan *Client
	Clients       map[*Client]bool
	ApiInfo       ApiInfo      // all api info
	UpdateMessage chan Message // updated info from api
}

type ApiInfo struct {
	sync.RWMutex
	Apis map[string]string
}

func (ai *ApiInfo) Update(msg Message) {
	ai.Lock()
	defer ai.Unlock()
	ai.Apis[msg.ApiType] = msg.Body
}

func NewManager() *Manager {
	return &Manager{
		Register:      make(chan *Client),
		Unregister:    make(chan *Client),
		Clients:       make(map[*Client]bool),
		ApiInfo:       ApiInfo{Apis: make(map[string]string)},
		UpdateMessage: make(chan Message),
	}
}

func (manager *Manager) Manage() {
	for {
		select {
		case client := <-manager.Register:
			manager.Clients[client] = true
			log.Printf("new client added. Num of clients: %v", len(manager.Clients))
			client.Conn.WriteJSON(manager.ApiInfo.Apis)
			break
		case client := <-manager.Unregister:
			//fmt.Println(client)
			delete(manager.Clients, client)
			break
		case message := <-manager.UpdateMessage:
			msg := make(map[string]string)
			manager.ApiInfo.Update(message)
			msg[message.ApiType] = message.Body
			for client, _ := range manager.Clients {
				client.Conn.WriteJSON(msg)
			}
			break
		}
	}
}
