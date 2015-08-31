package packet

import (
	. "gslib/utils"
	"math"
	"net"
)

const (
	PACKET_LIMIT = 65533 // 2^16 - 1 - 2
)

type Packet struct {
	pos  uint
	data []byte
}

func (p *Packet) Data() []byte {
	return p.data
}

func (p *Packet) Length() int {
	return len(p.data)
}

func (p *Packet) Pos() uint {
	return p.pos
}

func (p *Packet) Seek(n uint) {
	p.pos += n
}

//=============================================== Readers
func (p *Packet) ReadBool() (ret bool) {
	b := p.ReadByte()

	if b != byte(1) {
		return false
	}

	return true
}

func (p *Packet) ReadByte() (ret byte) {
	if p.pos >= uint(len(p.data)) {
		panic("read byte failed")
	}

	ret = p.data[p.pos]
	p.pos++
	return
}

func (p *Packet) ReadBytes() (ret []byte) {
	if p.pos+2 > uint(len(p.data)) {
		panic("read bytes header failed")
	}
	size := p.ReadUint16()
	if p.pos+uint(size) > uint(len(p.data)) {
		panic("read bytes data failed")
	}

	ret = p.data[p.pos : p.pos+uint(size)]
	p.pos += uint(size)
	return
}

func (p *Packet) ReadString() (ret string) {
	if p.pos+2 > uint(len(p.data)) {
		panic("read string header failed")
	}

	size := p.ReadUint16()
	if p.pos+uint(size) > uint(len(p.data)) {
		panic("read string data failed")
	}

	bytes := p.data[p.pos : p.pos+uint(size)]
	p.pos += uint(size)
	ret = string(bytes)
	return
}

func (p *Packet) ReadUint16() (ret uint16) {
	if p.pos+2 > uint(len(p.data)) {
		panic("read uint16 failed")
	}

	buf := p.data[p.pos : p.pos+2]
	ret = uint16(buf[0])<<8 | uint16(buf[1])
	p.pos += 2
	return
}

func (p *Packet) ReadInt16() (ret int16) {
	_ret := p.ReadUint16()
	ret = int16(_ret)
	return
}

func (p *Packet) ReadUint24() (ret uint32) {
	if p.pos+3 > uint(len(p.data)) {
		panic("read uint24 failed")
	}

	buf := p.data[p.pos : p.pos+3]
	ret = uint32(buf[0])<<16 | uint32(buf[1])<<8 | uint32(buf[2])
	p.pos += 3
	return
}

func (p *Packet) ReadInt24() (ret int32) {
	_ret := p.ReadUint24()
	ret = int32(_ret)
	return
}

func (p *Packet) ReadUint32() (ret uint32) {
	if p.pos+4 > uint(len(p.data)) {
		panic("read uint32 failed")
	}

	buf := p.data[p.pos : p.pos+4]
	ret = uint32(buf[0])<<24 | uint32(buf[1])<<16 | uint32(buf[2])<<8 | uint32(buf[3])
	p.pos += 4
	return
}

func (p *Packet) ReadInt32() (ret int32) {
	_ret := p.ReadUint32()
	ret = int32(_ret)
	return
}

func (p *Packet) ReadUint64() (ret uint64) {
	if p.pos+8 > uint(len(p.data)) {
		panic("read uint64 failed")
	}

	ret = 0
	buf := p.data[p.pos : p.pos+8]
	for i, v := range buf {
		ret |= uint64(v) << uint((7-i)*8)
	}
	p.pos += 8
	return
}

func (p *Packet) ReadInt64() (ret int64) {
	_ret := p.ReadInt64()
	ret = int64(_ret)
	return
}

func (p *Packet) ReadFloat32() (ret float32) {
	bits := p.ReadUint32()

	ret = math.Float32frombits(bits)
	if math.IsNaN(float64(ret)) || math.IsInf(float64(ret), 0) {
		return 0
	}

	return ret
}

func (p *Packet) ReadFloat64() (ret float64) {
	bits := p.ReadUint64()
	ret = math.Float64frombits(bits)
	if math.IsNaN(ret) || math.IsInf(ret, 0) {
		return 0
	}

	return ret
}

//================================================ Writers
func (p *Packet) WriteZeros(n int) {
	zeros := make([]byte, n)
	p.data = append(p.data, zeros...)
}

func (p *Packet) WriteBool(v bool) {
	if v {
		p.data = append(p.data, byte(1))
	} else {
		p.data = append(p.data, byte(0))
	}
}

func (p *Packet) WriteByte(v byte) {
	p.data = append(p.data, v)
}

func (p *Packet) WriteBytes(v []byte) {
	p.WriteUint16(uint16(len(v)))
	p.data = append(p.data, v...)
}

func (p *Packet) WriteRawBytes(v []byte) {
	p.data = append(p.data, v...)
}

func (p *Packet) WriteString(v string) {
	bytes := []byte(v)
	p.WriteUint16(uint16(len(bytes)))
	p.data = append(p.data, bytes...)
}

func (p *Packet) WriteUint16(v uint16) {
	buf := make([]byte, 2)
	buf[0] = byte(v >> 8)
	buf[1] = byte(v)
	p.data = append(p.data, buf...)
}

func (p *Packet) WriteInt16(v int16) {
	p.WriteUint16(uint16(v))
}

func (p *Packet) WriteUint24(v uint32) {
	buf := make([]byte, 3)
	buf[0] = byte(v >> 16)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v)
	p.data = append(p.data, buf...)
}

func (p *Packet) WriteUint32(v uint32) {
	buf := make([]byte, 4)
	buf[0] = byte(v >> 24)
	buf[1] = byte(v >> 16)
	buf[2] = byte(v >> 8)
	buf[3] = byte(v)
	p.data = append(p.data, buf...)
}

func (p *Packet) WriteInt32(v int32) {
	p.WriteUint32(uint32(v))
}

func (p *Packet) WriteUint64(v uint64) {
	buf := make([]byte, 8)
	for i := range buf {
		buf[i] = byte(v >> uint((7-i)*8))
	}

	p.data = append(p.data, buf...)
}

func (p *Packet) WriteInt64(v int64) {
	p.WriteUint64(uint64(v))
}

func (p *Packet) WriteFloat32(f float32) {
	v := math.Float32bits(f)
	p.WriteUint32(v)
}

func (p *Packet) WriteFloat64(f float64) {
	v := math.Float64bits(f)
	p.WriteUint64(v)
}

func Reader(data []byte) *Packet {
	return &Packet{pos: 0, data: data}
}

func Writer() *Packet {
	pkt := &Packet{pos: 0}
	pkt.data = make([]byte, 0, 128)
	return pkt
}

var encrypt = false

func (p *Packet) Send(conn net.Conn) {
	writer := Writer()
	if encrypt {
		data := Encrypt(p.Data())
		writer.WriteUint16(uint16(len(data)))
		writer.WriteRawBytes(data)
	} else {
		writer.WriteUint16(uint16(p.Length()))
		writer.WriteRawBytes(p.Data())
	}

	n, err := conn.Write(writer.Data())
	if err != nil {
		ERR("Error send reply, bytes:", n, "reason:", err)
		return
	}
}
