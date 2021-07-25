package incoming

import (
	"serapis/internal/pkg/network"
	"serapis/internal/pkg/protocol"
)

type Event interface {
	Handle(*network.Connection)
}

var Events = map[int]func(packet *protocol.Packet) Event{}
