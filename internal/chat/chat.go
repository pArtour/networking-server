package chat

import (
	"sync"

	"github.com/gofiber/websocket/v2"
)

type Connection struct {
	ConnID int64
	Conn   *websocket.Conn
}

var connections = make([]*Connection, 0)
var connMutex sync.Mutex

func AddConnection(ConnID int64, conn *websocket.Conn) {
	connMutex.Lock()
	connections = append(connections, &Connection{ConnID: ConnID, Conn: conn})
	connMutex.Unlock()
}

func RemoveConnection(connID int64) {
	connMutex.Lock()
	for i, conn := range connections {
		if conn.ConnID == connID {
			connections = append(connections[:i], connections[i+1:]...)
			break
		}
	}
	connMutex.Unlock()
}

func GetConnectionByID(ConnID int64) *Connection {
	for _, conn := range connections {
		if conn.ConnID == ConnID {
			return conn
		}
	}
	return nil
}

func BroadcastMessage(ConnID int64, message []byte) {
	for _, conn := range connections {
		if conn.ConnID == ConnID {
			conn.Conn.WriteMessage(websocket.TextMessage, message)
		}
	}
}
