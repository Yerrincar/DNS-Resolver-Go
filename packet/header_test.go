package packet

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeader(t *testing.T) {
	t.Run("Should encode a header into bytes", func(t *testing.T) {
		header := NewHeader(22, RECURSION_FLAG, 1, 0, 0, 0)

		encodeHeader := header.ToBytes()

		expected, err := hex.DecodeString("0016010000010000000000000")
		assert.NotNil(t, err)
		assert.Equal(t, expected, encodeHeader)
	})

	t.Run("Should create a header from a response", func(t *testing.T) {
		response, _ := hex.DecodeString("0016808000010002000000000")
		header, _ := ParseHeader(bytes.NewReader(response))

		assert.Equal(t, &Header{
			Id:      0x16,
			Flags:   1<<15 | 1<<7,
			QDCount: 0x1,
			ANCount: 0x2,
			NSCount: 0x0,
			ARCount: 0x0,
		}, header)
	})

	t.Run("Should return an error if the header flahs contain a query error", func(t *testing.T) {
		response, _ := hex.DecodeString("0016808000010002000000000")
		header, err := ParseHeader(bytes.NewReader(response))

		assert.Nil(t, header)
		assert.NotNil(t, err)
		assert.Equal(t, err, "error with the query")
	})
}
