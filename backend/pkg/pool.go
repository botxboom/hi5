package websocket

type Pool struct {
	Register chan *Client
	Unregister chan *Client
	Clients map[*Client]bool
	Broadcast chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register: make(chan *Client),
		Unregister: make(chan *Client),
		Clients: make(map[*Client]bool),
		Broadcast: make(chan Message),
	}
}

func (pool *Pool) Start(){
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			for client, _ := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: "New User joined..."})
			}
			break;
		
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			for client, _ := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: "User disconnected"})

			}
			break;
		case message := <-pool.Broadcast:
			for client, _ := range pool.Clients {
				if err := client.Conn.WriteJSON(message) 
				err != nil {
					return
				}
			}
		}
	}
}