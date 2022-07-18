package common

import (
	"fmt"
	"strings"
)

const EOFMarker = "%%END%%"

func AssertNoError(err error) {
	if err != nil {
		panic(err)
	}
}

func AssertNotEmpty(text string, message string) {
	if strings.Trim(text, " ") == "" {
		fmt.Println(message)

		panic("Unexpected empty field error")
	}
}

func ConvertText2Bits(text string) []byte {
	data := []byte(text)

	var bitArr []byte
	for i := range data {
		for j := 0; j < 8; j++ {
			b := (data[i] & (1 << j) >> j)

			bitArr = append(bitArr, b)
		}
	}
	return bitArr
}

func ConvertBitArray2Text(data []byte) string {
	length := len(data)
	length = length - length%8

	var bytes []byte
	for i := 0; i < length; i += 8 {
		var b uint8
		for j := 0; j < 8; j++ {
			b += data[i+j] << (j)
		}
		bytes = append(bytes, b)
	}
	return string(bytes)
}
