package network

import (
	"github.com/gorilla/websocket"

	"serapis/internal/pkg/messages/outgoing"
)

type Connection struct {
	ws *websocket.Conn
}

func NewConnection(ws *websocket.Conn) *Connection {
	return &Connection{ws: ws}
}

func (c *Connection) Write(composer outgoing.Composer) error {
	return c.ws.WriteMessage(websocket.BinaryMessage, composer.Bytes())
}

func (c *Connection) Read() ([]byte, error) {
	_, data, err := c.ws.ReadMessage()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Connection) Close() error {
	return c.ws.Close()
}
