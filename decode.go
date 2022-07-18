package steganogopher

import (
	"image"
	png "image/png"
	"os"
	"steganogopher/internal/common"
	"strings"
)

// Decode is used to extract/decode a hidden text message from a .PNG image
func Decode(inputPath string) string {
	inputPath = strings.Trim(inputPath, " ")
	common.AssertNotEmpty(inputPath, "-i=<PATH-TO-IMAGE> should be not empty")

	data := extractHiddenDataFromFile(inputPath)

	decoded := common.ConvertBitArray2Text(data)

	return decoded[:len(decoded)-len(common.EOFMarker)]
}

func extractHiddenDataFromFile(path string) []byte {
	fIn, err := os.Open(path)
	common.AssertNoError(err)
	defer fIn.Close()

	imgData, err := png.Decode(fIn)
	common.AssertNoError(err)

	img := imgData.(*image.RGBA)

	maxBits := int(img.Rect.Max.Y * img.Rect.Max.X * 3)
	current := 0

	data := []byte{}
	for y := 0; y < img.Rect.Max.Y && current < maxBits; y++ {
		for x := 0; x < img.Rect.Max.X && current < maxBits; x++ {
			r, g, b, _ := img.At(x, y).RGBA()

			for c := 0; c < 3 && current < maxBits && !checkEOFMarker(data); c++ {
				current++

				switch c {
				case 0:
					data = append(data, uint8(r)&1)
				case 1:
					data = append(data, uint8(g)&1)
				case 2:
					data = append(data, uint8(b)&1)
				}
			}
		}
	}
	return data
}

func checkEOFMarker(data []byte) bool {
	encodedMarker := common.ConvertText2Bits(common.EOFMarker)
	markerLen := len(encodedMarker)

	if len(data) < markerLen {
		return false
	}
	data = data[len(data)-markerLen:]
	for i := 0; i < markerLen; i++ {
		if data[i] != encodedMarker[i] {
			return false
		}
	}
	return true
}
