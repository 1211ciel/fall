package internal

import "github.com/1211ciel/fall/im/comet/internal/svc"

var GlobalHub *Hub

// Hub maintains the set of active clients and broadcasts messages to the clients.
type Hub struct {
	// 服务器地址 etcd 存放注册在etcd上面的 地址
	ServerAddr string
	// Registered clients.
	clients map[*Channel]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Channel

	// Unregister requests from clients.
	unregister chan *Channel

	// 具备远程调用能力
	svc *svc.ServiceContext
}

func NewHub(svc *svc.ServiceContext) *Hub {
	hub := Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Channel),
		unregister: make(chan *Channel),
		clients:    make(map[*Channel]bool),
		svc:        svc,
	}
	GlobalHub = &hub
	return &hub
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register: // 注册用户连接
			h.clients[client] = true
			println(len(h.clients), h.clients)
		case client := <-h.unregister: // 注销用户连接
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast: // 广播给当前服务器的所有用户
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
