package main

import (
	"flag"
	"fmt"
	st "steganogopher"
)

func main() {
	var flagEncode, flagDecode bool
	var inputPath, outputPath, message string

	flag.BoolVar(&flagDecode, "d", false, "decode data from a .PNG file")
	flag.BoolVar(&flagEncode, "e", false, "encode data into a .PNG file")
	flag.StringVar(&inputPath, "i", "", "path to the input .PNG file")
	flag.StringVar(&outputPath, "o", "", "path to the output .PNG file. Default value is <INPUT>_X.PNG")
	flag.StringVar(&message, "m", "", "message to be encoded")
	flag.Parse()

	fmt.Printf("==| SteganoGophy |==\n\n")

	if flagEncode {
		st.Encode(message, inputPath, outputPath)
	} else if flagDecode {
		msg := st.Decode(inputPath)

		fmt.Printf("Message: \"%s\"\n", msg)
	}
}
