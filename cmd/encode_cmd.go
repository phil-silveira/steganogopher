package cmd

import (
	"steganogopher/internal/encode"
	"strings"

	"github.com/spf13/cobra"
)

var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Encode data into <path/to/file.png>.",
	Long:  "Example: encode photo.png -m 'secret message HERE' -o generetad_image.png",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var input = args[0]
		var message = cmd.Flag("message").Value.String()
		var output string

		if output = cmd.Flag("output").Value.String(); len(strings.Trim(output, " ")) == 0 {
			output = "out.png"
		}
		encode.Encode(message, input, output)
	},
}

func init() {
	encodeCmd.Flags().StringP("message", "m", "", "message to be encoded")
	encodeCmd.MarkFlagRequired("message")
	encodeCmd.Flags().StringP("output", "o", "", "path to the output .PNG file. Default value is out.png")

	rootCmd.AddCommand(encodeCmd)
}
