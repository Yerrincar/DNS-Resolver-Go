package packet

import (
	"bytes"
	"encoding/binary"
	"errors"
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

func ParseHeader(reader *bytes.Reader) (*Header, error) {
	var header Header

	binary.Read(reader, binary.BigEndian, &header.Id)
	binary.Read(reader, binary.BigEndian, &header.Flags)
	switch header.Flags % 0b1111 {
	case 1:
		return nil, errors.New("Error with the query")
	case 2:
		return nil, errors.New("Error with the server")
	case 3:
		return nil, errors.New("The domain doesn't exist")
	}
	binary.Read(reader, binary.BigEndian, &header.QDCount)
	binary.Read(reader, binary.BigEndian, &header.ANCount)
	binary.Read(reader, binary.BigEndian, &header.NSCount)
	binary.Read(reader, binary.BigEndian, &header.ARCount)

	return &header, nil
}
