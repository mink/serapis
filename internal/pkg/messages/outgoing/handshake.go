package outgoing

import "serapis/internal/pkg/protocol"

type SecureLoginOKComposer struct {
	header uint16
}

func (c *SecureLoginOKComposer) Bytes() []byte {
	return protocol.NewOutgoingPacket(c.header).Bytes()
}

func NewSecureLoginOKComposer() *SecureLoginOKComposer {
	return &SecureLoginOKComposer{header: 2491}
}
