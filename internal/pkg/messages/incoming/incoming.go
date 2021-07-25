package incoming

import (
	"github.com/gorilla/websocket"

	"serapis/internal/pkg/protocol"
)

type Event interface {
	Handle(*websocket.Conn)
}

var Events = map[int]func(packet *protocol.Packet) Event{}
