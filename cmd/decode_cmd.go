package cmd

import (
	"steganogopher/internal/common"
	"steganogopher/internal/decode"
	"strings"

	"github.com/spf13/cobra"
)

var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decode data from <path/to/file.png>.",
	Long:  "Example: decode photo.png  -o decoded_message.txt",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var input = args[0]
		var output string

		if output = cmd.Flag("output").Value.String(); len(strings.Trim(output, " ")) == 0 {
			output = "out.png"
		}
		message := decode.Decode(input)

		common.WriteStringIntoFile(output, message)
	},
}

func init() {
	decodeCmd.Flags().StringP("output", "o", "", "path to the output .TXT file. Default value is out.txt")

	rootCmd.AddCommand(decodeCmd)
}
