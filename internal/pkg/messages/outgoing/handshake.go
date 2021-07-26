package outgoing

import "serapis/internal/pkg/protocol"

type SecureLoginOKComposer struct {
	header uint16
}

func (c *SecureLoginOKComposer) Bytes() []byte {
	return protocol.NewOutgoingPacket(2491).Bytes()
}
