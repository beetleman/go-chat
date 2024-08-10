package encoding_test

import (
	"testing"

	"github.com/beetleman/go-chat/internal/encoding"
	"github.com/stretchr/testify/assert"
)

func TestEncodeDecode(t *testing.T) {
	for _, text := range []string{"Hello", "🐐🩷", "A=πr²"} {
		message := encoding.Message{Text: text}
		assert.Equal(t, text, encoding.Decode(message.Encode()).Text)
	}
}
