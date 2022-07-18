package steganogopher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeTextFromPNG(t *testing.T) {
	assert := assert.New(t)
	msg := Decode("./examples/avatar_test.png")

	assert.Equal("hidden message XD", msg)

}
