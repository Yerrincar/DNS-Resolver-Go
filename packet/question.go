package packet

import (
	"bytes"
	"encoding/binary"
)

const TYPE_A uint16 = 1
const TYPE_NS uint16 = 2
const CLASS_IN uint16 = 1

type Question struct {
	Qname  []byte
	Qtype  uint16
	Qclass uint16
}

func NewQuestion(qname string, qtype, qclass uint16) *Question {
	return &Question{
		Qname:  encodeDnsName([]byte(qname)),
		Qtype:  qtype,
		Qclass: qclass,
	}
}

func (q *Question) ToBytes() []byte {
	encodedQuestion := new(bytes.Buffer)
	binary.Write(encodedQuestion, binary.BigEndian, q.Qname)
	binary.Write(encodedQuestion, binary.BigEndian, q.Qtype)
	binary.Write(encodedQuestion, binary.BigEndian, q.Qclass)

	return encodedQuestion.Bytes()
}

func encodeDnsName(qname []byte) []byte {
	var encoded []byte
	parts := bytes.Split([]byte(qname), []byte{'.'})
	for _, part := range parts {
		encoded = append(encoded, byte(len(part)))
		encoded = append(encoded, part...)
	}

	return append(encoded, 0x00)

}
