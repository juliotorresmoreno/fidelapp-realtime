package ws

type Hub struct {
	clients          map[string][]*UserClient
	Register         chan *UserClient
	Unregister       chan string
	unregisterClient chan *UserClient
}

func NewHub() *Hub {
	return &Hub{
		Register:   make(chan *UserClient),
		Unregister: make(chan string),
		clients:    map[string][]*UserClient{},
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			if _, ok := h.clients[client.id]; !ok {
				h.clients[client.id] = make([]*UserClient, 0)
			}
			h.clients[client.id] = append(h.clients[client.id], client)
		case clientId := <-h.Unregister:
			if _, ok := h.clients[clientId]; ok {
				for _, client := range h.clients[clientId] {
					close(client.send)
				}
				h.clients[clientId] = make([]*UserClient, 0)
			}
		}
	}
}
