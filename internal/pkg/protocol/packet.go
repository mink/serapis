package protocol

import "encoding/binary"

type Packet struct {
	header uint16
	data   []byte
}

func NewInboundPacket(data []byte) *Packet {
	p := &Packet{data: data}
	p.readLength()
	p.header = p.readHeader()
	return p
}

func NewOutgoingPacket(header uint16) *Packet {
	return &Packet{header: header}
}

func (p *Packet) pop(bytes int) {
	p.data = p.data[bytes:]
}

func (p *Packet) readLength() uint32 {
	length := binary.BigEndian.Uint32(p.data[0:4])
	p.pop(4)
	return length
}

func (p *Packet) readHeader() uint16 {
	header := binary.BigEndian.Uint16(p.data[0:2])
	p.pop(2)
	return header
}

func (p *Packet) ReadString() string {
	length := binary.BigEndian.Uint16(p.data[0:2])
	p.pop(2)
	str := p.data[0:length]
	p.pop(int(length))
	return string(str)
}

func (p Packet) Header() uint16 {
	return p.header
}

func (p Packet) Length() uint32 {
	return uint32(len(p.Data()) + 2)
}

func (p Packet) Data() []byte {
	return p.data
}

func (p Packet) Bytes() []byte {
	var bytes []byte

	lengthBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBytes, p.Length())
	bytes = append(bytes, lengthBytes...)

	headerBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(headerBytes, p.header)
	bytes = append(bytes, headerBytes...)

	bytes = append(bytes, p.data...)

	return bytes
}
