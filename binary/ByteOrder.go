package binary

import "encoding/binary"

type ByteOrder interface {
	Uint16([]byte) uint16
	Uint32([]byte) uint32
	Uint64([]byte) uint64
	PutUint16([]byte, uint16)
	PutUint32([]byte, uint32)
	PutUint64([]byte, uint64)
}
type bigEndian struct{}

func (bigEndian) Uint16(b []byte) uint16 {
	if Reference {
		return binary.BigEndian.Uint16(b)
	}
	return uint16(b[0])<<8 | uint16(b[1])
}
func (bigEndian) Uint32(b []byte) uint32 {
	if Reference {
		return binary.BigEndian.Uint32(b)
	}
	return uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
}
func (bigEndian) Uint64(b []byte) uint64 {
	if Reference {
		return binary.BigEndian.Uint64(b)
	}
	return uint64(b[0])<<56 | uint64(b[1])<<48 | uint64(b[2])<<40 | uint64(b[3])<<32 | uint64(b[4])<<24 | uint64(b[5])<<16 | uint64(b[6])<<8 | uint64(b[7])
}
func (bigEndian) PutUint16(b []byte, v uint16) {
	if Reference {
		binary.BigEndian.PutUint16(b, v)
		return
	}
	b[0], b[1] = byte(v>>8), byte(v)
}
func (bigEndian) PutUint32(b []byte, v uint32) {
	if Reference {
		binary.BigEndian.PutUint32(b, v)
		return
	}
	b[0], b[1], b[2], b[3] = byte(v>>24), byte(v>>16), byte(v>>8), byte(v)
}
func (bigEndian) PutUint64(b []byte, v uint64) {
	if Reference {
		binary.BigEndian.PutUint64(b, v)
		return
	}
	b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7] = byte(v>>56), byte(v>>48), byte(v>>40), byte(v>>32), byte(v>>24), byte(v>>16), byte(v>>8), byte(v)
}
