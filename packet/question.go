package packet

import (
	"bytes"
	"encoding/binary"
	"io"
	"strings"
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

func ParseQuestion(reader *bytes.Reader) *Question {
	var question Question

	question.Qname = []byte(DecodeName(reader))
	binary.Read(reader, binary.BigEndian, &question.Qtype)
	binary.Read(reader, binary.BigEndian, &question.Qclass)

	return &question
}

func DecodeName(reader *bytes.Reader) string {
	var name bytes.Buffer

	for {
		lengthByte, _ := reader.ReadByte()

		if (lengthByte & 0xC0) == 0xC0 {
			name.WriteString(getBackTheDomainFromHeader(reader, lengthByte))
			break
		}

		if lengthByte == 0 {
			break
		}

		label := make([]byte, lengthByte)
		io.ReadFull(reader, label)
		name.Write(label)
		name.WriteByte('.')
	}

	result, _ := strings.CutSuffix(name.String(), ".")
	return result
}

func getBackTheDomainFromHeader(reader *bytes.Reader, lengthByte byte) string {
	nextByte, _ := reader.ReadByte()
	pointer := uint16(uint16(lengthByte)&0x3F) | uint16(nextByte)

	currentPos, _ := reader.Seek(0, io.SeekCurrent)

	reader.Seek(int64(pointer), io.SeekStart)

	decodedName := DecodeName(reader)

	reader.Seek(currentPos, io.SeekStart)

	return decodedName
}
