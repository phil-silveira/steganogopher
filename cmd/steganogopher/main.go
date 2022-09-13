package main

import (
	"flag"
	"fmt"
	"strings"

	st "steganogopher"
)

var commands = map[string]string{
	"d": "decode data from <path/to/file.png>",
	"e": "encode data into <path/to/file.png>",
	"o": "path to the output .PNG file. Default value is <INPUT>_X.PNG",
	"m": "message to be encoded",
}

func main() {
	var encodeTarget, decodeTarget string
	var outputPath, message string

	flag.StringVar(&decodeTarget, "d", "", commands["d"])
	flag.StringVar(&encodeTarget, "e", "", commands["e"])
	flag.StringVar(&outputPath, "o", "", commands["o"])
	flag.StringVar(&message, "m", "", commands["m"])
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

	fmt.Printf("Steganogopher is a steganography tool to play with .PNG images.\n\nUsage:\n\n\tsteganogopher <command> [arguments]\n\nThe commands are:\n\n")
	for k, v := range commands {
		fmt.Printf("\t-%s \t%s\n", k, v)
	}
}
