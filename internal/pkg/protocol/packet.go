package protocol

import "encoding/binary"

type Packet struct {
	header uint16
	length uint32
	data   []byte
}

func NewPacket(data []byte) *Packet {
	p := &Packet{data: data}
	p.length = p.readLength()
	p.header = p.readHeader()
	return p
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
	return p.length
}

func (p Packet) Data() []byte {
	return p.data
}
