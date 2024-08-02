package ws

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type OrderStatusUpdate struct {
	OrderID string `json:"order_id"`
	Status  string `json:"status"`
}

type WebSocketHub struct {
	clients    map[string]map[*websocket.Conn]bool // map[orderID]map[conn]bool
	broadcast  chan OrderStatusUpdate
	register   chan subscription
	unregister chan subscription
	mu         sync.Mutex
}

type subscription struct {
	conn    *websocket.Conn
	orderID string
}

var hub = WebSocketHub{
	clients:    make(map[string]map[*websocket.Conn]bool),
	broadcast:  make(chan OrderStatusUpdate),
	register:   make(chan subscription),
	unregister: make(chan subscription),
}

func (h *WebSocketHub) run() {
	for {
		select {
		case sub := <-h.register:
			h.mu.Lock()
			if _, ok := h.clients[sub.orderID]; !ok {
				h.clients[sub.orderID] = make(map[*websocket.Conn]bool)
			}
			h.clients[sub.orderID][sub.conn] = true
			h.mu.Unlock()
		case sub := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[sub.orderID][sub.conn]; ok {
				delete(h.clients[sub.orderID], sub.conn)
				sub.conn.Close()
				if len(h.clients[sub.orderID]) == 0 {
					delete(h.clients, sub.orderID)
				}
			}
			h.mu.Unlock()
		case update := <-h.broadcast:
			h.mu.Lock()
			if clients, ok := h.clients[update.OrderID]; ok {
				for conn := range clients {
					err := conn.WriteJSON(update)
					if err != nil {
						conn.Close()
						delete(clients, conn)
					}
				}
				if len(clients) == 0 {
					delete(h.clients, update.OrderID)
				}
			}
			h.mu.Unlock()
		}
	}
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	orderID := r.URL.Query().Get("order_id")
	if orderID == "" {
		http.Error(w, "Missing order_id", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %s", err)
		return
	}

	sub := subscription{conn: conn, orderID: orderID}
	hub.register <- sub

	defer func() {
		hub.unregister <- sub
	}()

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}
}

func BroadcastOrderStatusUpdate(update OrderStatusUpdate) {
	hub.broadcast <- update
}

func StartWebSocketHub() {
	go hub.run()
}
