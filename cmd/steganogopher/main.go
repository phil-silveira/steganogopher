package main

import (
	"flag"
	"fmt"
	st "steganogopher"
	"strings"
)

func main() {
	var encodeTarget, decodeTarget string
	var outputPath, message string

	flag.StringVar(&decodeTarget, "d", "", "decode data from <path/to/file.png>")
	flag.StringVar(&encodeTarget, "e", "", "encode data into <path/to/file.png>")
	flag.StringVar(&outputPath, "o", "", "path to the output .PNG file. Default value is <INPUT>_X.PNG")
	flag.StringVar(&message, "m", "", "message to be encoded")
	flag.Parse()

	decodeTarget = strings.Trim(decodeTarget, " ")
	encodeTarget = strings.Trim(encodeTarget, " ")
	outputPath = strings.Trim(outputPath, " ")
	message = strings.Trim(message, " ")

	if encodeTarget != "" {
		st.Encode(message, encodeTarget, outputPath)
		return
	}

	if decodeTarget != "" {
		msg := st.Decode(decodeTarget)

		fmt.Printf("{\"message\": \"%s\"}\n", msg)
		return
	}
}
