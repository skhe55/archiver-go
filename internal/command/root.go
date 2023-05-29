package command

import (
	"log"

	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Short: "Archiver",
}

func Execute() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
