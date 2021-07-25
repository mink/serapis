package incoming

import (
	"serapis/internal/pkg/protocol"
)

type Event interface {
	Handle()
}

var Events = map[int]func(packet *protocol.Packet) Event{}
