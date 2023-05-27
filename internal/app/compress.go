package app

import "github.com/spf13/cobra"

var compressCmd = &cobra.Command{
	Use:   "compress",
	Short: "Compress file",
}

func init() {
	cmd.AddCommand(compressCmd)
}
