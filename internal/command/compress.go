package command

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/skhe55/archiver-go/internal/services/compression"
	"github.com/skhe55/archiver-go/internal/services/compression/vlc"
	"github.com/spf13/cobra"
)

var compressCmd = &cobra.Command{
	Use:   "compress",
	Short: "Compress file",
	Run:   compress,
}

const compressedExtension = "vlc"

func compress(cmd *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		log.Fatal("path to file is not provided")
	}
	var encoder compression.Encoder

	method := cmd.Flag("method").Value.String()

	switch method {
	case "vlc":
		encoder = vlc.New()
	default:
		cmd.PrintErr("unknown method...")
	}
	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	compressed := encoder.Encode(string(data))

	err = os.WriteFile(compressedFileName(filePath), compressed, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func compressedFileName(path string) string {
	fName := filepath.Base(path)
	return strings.TrimSuffix(fName, filepath.Ext(fName)) + "." + compressedExtension
}

func init() {
	cmd.AddCommand(compressCmd)

	compressCmd.Flags().StringP("method", "m", "", "compression method: vlc, etc")

	if err := compressCmd.MarkFlagRequired("method"); err != nil {
		panic("flag m or 'method' not provided " + err.Error())
	}
}
