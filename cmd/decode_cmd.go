package cmd

import (
	"fmt"
	"steganogopher/internal/common"
	"steganogopher/internal/decode"

	"github.com/spf13/cobra"
)

var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decode data from <path/to/file.png>.",
	Long: `Decode data from <path/to/file.png>. 
		Example: decode photo.png  -o decoded_message.txt`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var input = args[0]

		message := decode.Decode(input)

		outputFilename := cmd.Flag("output").Value.String()

		hasNoOutputFile := len(outputFilename) == 0
		if hasNoOutputFile {
			printDecodedMessage(message)
			return
		}
		common.WriteStringIntoFile(outputFilename, message)
	},
}

func printDecodedMessage(message string) {
	fmt.Printf(`(decoded): %s`, message)
}

func init() {
	decodeCmd.Flags().StringP("output", "o", "", "path to the output .TXT file. Default value is out.txt")

	rootCmd.AddCommand(decodeCmd)
}
