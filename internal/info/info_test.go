package info

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetImageInfoShouldReturn468X698ImageInfo(t *testing.T) {
	res, err := GetImageInfo("../../examples/le-petit-prince.png")

	assert.NoError(t, err)
	assert.Equal(t, "le-petit-prince.png", res.Name)
	assert.Equal(t, "PNG", res.DocumentType)
	assert.Equal(t, 468, res.Width)
	assert.Equal(t, 698, res.Height)
	assert.Equal(t, 122499, res.AvailableSpaceInBytes)
	assert.Equal(t, "119 KB", res.AvailableSpaceFormatted)
}

func TestFormatSpaceShouldReturnBytes(t *testing.T) {
	res := formatSpace(1023)

	assert.Equal(t, "1023 B", res)
}

func TestFormatSpaceShouldReturnKBytes(t *testing.T) {
	res := formatSpace(1025)

	assert.Equal(t, "1 KB", res)
}
