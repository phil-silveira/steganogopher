package steganogopher

import (
	"fmt"
	"image"
	"image/color"
	png "image/png"
	"os"
	"steganogopher/internal/common"
	"strings"
)

// Encode is used to encode/hide a text message into a .PNG image
func Encode(text string, inputPath string, outputPath string) {
	text = strings.Trim(text, " ")
	common.AssertNotEmpty(text, "-m=<text> should be not empty")
	text += common.EOFMarker

	inputPath = strings.Trim(inputPath, " ")
	common.AssertNotEmpty(inputPath, "-i=<PATH-TO-IMAGE> should be not empty")

	outputPath = strings.Trim(outputPath, " ")
	if outputPath == "" {
		outputPath = inputPath[:len(inputPath)-4] + ".out" + inputPath[len(inputPath)-4:]
		fmt.Printf("***Warning: -o=<OUTPUT-PATH> is empty, default value is '%s'.\n", outputPath)
	}

	bits := common.ConvertText2Bits(text)

	hideDataIntoFile(bits, inputPath, outputPath)
}

func hideDataIntoFile(data []byte, inputPath string, outputPath string) {
	fIn, err := os.Open(inputPath)
	common.AssertNoError(err)
	defer fIn.Close()

	fOut, err := os.Create(outputPath)
	common.AssertNoError(err)
	defer fOut.Close()

	imgData, err := png.Decode(fIn)
	common.AssertNoError(err)

	img := imgData.(*image.NRGBA)

	count := len(data)
	current := 0
	for y := 0; y < img.Rect.Max.Y && current < count; y++ {
		for x := 0; x < img.Rect.Max.X && current < count; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			var nR, nG, nB uint8

			for c := 0; c < 3 && current < count; c++ {
				switch c {
				case 0:
					nR = data[current]
				case 1:
					nG = data[current]
				case 2:
					nB = data[current]
				}
				current++
			}

			img.Set(
				x, y,
				color.RGBA{
					R: ((uint8(r) >> 1) << 1) | nR,
					G: ((uint8(g) >> 1) << 1) | nG,
					B: ((uint8(b) >> 1) << 1) | nB,
					A: uint8(a),
				},
			)
		}
	}
	png.Encode(fOut, imgData)
}
