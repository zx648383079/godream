package server

import (
	"github.com/gin-gonic/gin"
	"zodream.cn/godream/modules/chat/models"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan *models.MessageItem

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

// NewHub .
func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan *models.MessageItem),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

// Run .
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true // 注册client端
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			msgByte := encodeJSON(EVENT_MESSAGE, gin.H{
				"data": message,
			})
			for client := range h.clients {
				if !message.IsRoom(client.roomType, client.roomId) {
					continue
				}
				client.send <- msgByte
			}
		}
	}
}
