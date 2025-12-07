package packet

import (
	"bytes"
	"encoding/binary"
)

const RECURSION_FLAG uint16 = 1 << 8

type Header struct {
	Id      uint16
	Flags   uint16
	QDCount uint16
	ANCount uint16
	NSCount uint16
	ARCount uint16
}

func NewHeader(id, flags, qdcount, ancount, nscount, arcount uint16) *Header {
	return &Header{
		Id:      id,
		Flags:   flags,
		QDCount: qdcount,
		ANCount: ancount,
		NSCount: nscount,
		ARCount: arcount,
	}
}

func (header *Header) ToBytes() []byte {
	encodeHeader := new(bytes.Buffer)
	binary.Write(encodeHeader, binary.BigEndian, header.Id)
	binary.Write(encodeHeader, binary.BigEndian, header.Flags)
	binary.Write(encodeHeader, binary.BigEndian, header.QDCount)
	binary.Write(encodeHeader, binary.BigEndian, header.ANCount)
	binary.Write(encodeHeader, binary.BigEndian, header.NSCount)
	binary.Write(encodeHeader, binary.BigEndian, header.ARCount)

	return encodeHeader.Bytes()
}
