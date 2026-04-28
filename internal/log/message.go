package log

import "encoding/binary"

type Message struct {
	Offset int64
	Timestamp int64
	Payload []byte
}

func (m Message) Encode() []byte {
	buf := make([]byte, 4+8+8+len(m.Payload))
	binary.BigEndian.PutUint32(buf[0:4], uint32(len(m.Payload)))
	binary.BigEndian.PutUint64(buf[4:12], uint64(m.Offset))
	binary.BigEndian.PutUint64(buf[12:20], uint64(m.Timestamp))
	copy(buf[20:], m.Payload)

	return buf
}

func Decode(buf []byte) Message {
	bodyLen := int(binary.BigEndian.Uint32(buf[0:4]))
	return Message{
		Offset: int64(binary.BigEndian.Uint64(buf[4:12])),
		Timestamp: int64(binary.BigEndian.Uint64(buf[12:20])),
		Payload: buf[20: 20+bodyLen],
	}
}