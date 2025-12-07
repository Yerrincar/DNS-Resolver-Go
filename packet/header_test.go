package packet

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeader(t *testing.T) {
	t.Run("Should encode a header into bytes", func(t *testing.T) {
		header := NewHeader(22, RECURSION_FLAG, 1, 0, 0, 0)

		encodeHeader := header.ToBytes()

		expected, err := hex.DecodeString("0016010000010000000000000")
		assert.NotNil(t, err)
		assert.Equal(t, expected, encodeHeader)
	})
}
