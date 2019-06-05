package views

import (
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
	log.Println("Client connected!")
	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		log.Printf("Client sent %d %s %s\n", msgType, msg, r.RemoteAddr)
	}
}
