package chat

import (
	"sync"

	"github.com/gofiber/websocket/v2"
)

type Connection struct {
	UserID int64
	Conn   *websocket.Conn
}

var connections = make([]*Connection, 0)
var connMutex sync.Mutex

func AddConnection(userID int64, conn *websocket.Conn) {
	connMutex.Lock()
	connections = append(connections, &Connection{UserID: userID, Conn: conn})
	connMutex.Unlock()
}

func RemoveConnection(userID int64) {
	connMutex.Lock()
	for i, conn := range connections {
		if conn.UserID == userID {
			connections = append(connections[:i], connections[i+1:]...)
			break
		}
	}
	connMutex.Unlock()
}

func GetConnectionByUserID(userID int64) *Connection {
	for _, conn := range connections {
		if conn.UserID == userID {
			return conn
		}
	}
	return nil
}

func BroadcastMessage(senderID int64, message string) {
	for _, conn := range connections {
		if conn.UserID != senderID {
			conn.Conn.WriteMessage(websocket.TextMessage, []byte(message))
		}
	}
}
