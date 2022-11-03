package cmd

import (
	"fmt"
	"steganogopher/internal/info"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Display information about the specified file.",
	Long: `Display information about the specified file.
		Example: info photo.png`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var inputImagePath = args[0]

		res, _ := info.GetImageInfo(inputImagePath)

		fmt.Printf(`
Info:
	File name:  %s
	Document type:  %s
	Image size:  %d x %d pixels

	Available space:  %s
`,
			res.Name,
			res.DocumentType,
			res.Width,
			res.Height,
			res.AvailableSpaceFormatted,
		)
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
