package views

import (
	"fmt"
	"hub"
	"log"
	"mux"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func init() {
	mux.GetMux().HandleFunc("/wsChannel", wsHandler)
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	defer conn.Close()
	if err != nil {
		log.Println(err)
		return
	}
	hub.AddConnection(conn)
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			//client probably disconnected
			return
		}
		hub.Broadcast(fmt.Sprintf("%s: %s", msg, r.RemoteAddr))
	}
}
