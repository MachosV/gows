package hub

import (
	"log"

	"github.com/gorilla/websocket"
)

var connections []*websocket.Conn

func init() {
	log.Println("Hub online")
}

func AddConnection(conn *websocket.Conn) int {
	connections = append(connections, conn)
	return len(connections) - 1 //return the index to the caller, to be able to remove it later
}

func RemoveConnection(index int) {
	connections[index] = connections[len(connections)-1]
	connections = connections[:len(connections)-1]
}

func Broadcast(msg string) {
	for index := range connections {
		connections[index].WriteMessage(websocket.TextMessage, []byte(msg))
	}
}
