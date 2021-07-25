package protocol

import "encoding/binary"

type packet struct {
	header uint16
	length uint32
	data   []byte
}

func NewPacket(data []byte) *packet {
	p := &packet{data: data}
	p.length = p.readLength()
	p.header = p.readHeader()
	return p
}

func (p *packet) pop(bytes int) {
	p.data = p.data[bytes:]
}

func (p *packet) readLength() uint32 {
	length := binary.BigEndian.Uint32(p.data[0:4])
	p.pop(4)
	return length
}

func (p *packet) readHeader() uint16 {
	header := binary.BigEndian.Uint16(p.data[0:2])
	p.pop(2)
	return header
}

func (p *packet) ReadString() string {
	length := binary.BigEndian.Uint16(p.data[0:2])
	p.pop(2)
	str := p.data[0:length]
	p.pop(int(length))
	return string(str)
}

func (p packet) Header() uint16 {
	return p.header
}

func (p packet) Length() uint32 {
	return p.length
}

func (p packet) Data() []byte {
	return p.data
}
