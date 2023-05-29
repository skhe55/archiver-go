package command

import (
	"github.com/spf13/cobra"
)

var uncompressCmd = &cobra.Command{
	Use:   "uncompress",
	Short: "Uncompress file",
}

func init() {
	cmd.AddCommand(uncompressCmd)
}
