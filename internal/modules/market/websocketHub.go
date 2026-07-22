package market

import (
	"encoding/json"
	"log"
	"strings"
	"sync"

	"github.com/gofiber/contrib/websocket"
)

type Hub struct {
	Rooms  map[string]map[*websocket.Conn]struct{}
	Global map[*websocket.Conn]struct{}
	mu     sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		Rooms:  make(map[string]map[*websocket.Conn]struct{}),
		Global: make(map[*websocket.Conn]struct{}),
	}
}

func (h *Hub) Add(symbol string, conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()

	symbol = strings.ToUpper(symbol)

	if h.Rooms[symbol] == nil {
		h.Rooms[symbol] = make(map[*websocket.Conn]struct{})
	}

	h.Rooms[symbol][conn] = struct{}{}
}

func (h *Hub) Remove(symbol string, conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()

	symbol = strings.ToUpper(symbol)

	if room, ok := h.Rooms[symbol]; ok {
		delete(room, conn)

		if len(room) == 0 {
			delete(h.Rooms, symbol)
		}
	}

	conn.Close()
}

func (h *Hub) Broadcast(symbol string, payload any) {
	h.mu.RLock()
	room, ok := h.Rooms[strings.ToUpper(symbol)]
	if !ok {
		h.mu.RUnlock()
		return
	}

	conns := make([]*websocket.Conn, 0, len(room))
	for c := range room {
		conns = append(conns, c)
	}
	h.mu.RUnlock()

	b, _ := json.Marshal(payload)

	for _, c := range conns {
		if err := c.WriteMessage(websocket.TextMessage, b); err != nil {
			go h.Remove(symbol, c)
		}
	}
}

func (h *Hub) AddGlobal(conn *websocket.Conn) {

	h.mu.Lock()
	defer h.mu.Unlock()

	h.Global[conn] = struct{}{}

	log.Println("Global clients:",len(h.Global))
}

func (h *Hub) RemoveGlobal(conn *websocket.Conn) {

	h.mu.Lock()
	defer h.mu.Unlock()

	delete(h.Global,conn)

	// conn.Close()
}

func (h *Hub) BroadcastGlobal(payload any) {

	h.mu.RLock()

	log.Println(
		"Global clients:",
		len(h.Global),
	)

	conns := make([]*websocket.Conn,0,len(h.Global))

	for c := range h.Global {

		conns = append(conns,c)
	}

	h.mu.RUnlock()

	data, err :=json.Marshal(payload)
	if err != nil {

		log.Println(
			"marshal error:",
			err,
		)

		return
	}

	for _,c:= range conns {

		err:=c.WriteMessage(websocket.TextMessage,data)

		if err!=nil {
			log.Println(
				"write error:",
				err,
			)
			go h.RemoveGlobal(c)
		}
	}
}
