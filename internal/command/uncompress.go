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

var uncompressCmd = &cobra.Command{
	Use:   "uncompress",
	Short: "Uncompress file",
	Run:   uncompress,
}

const uncompressedExtension = "txt"

func uncompress(cmd *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		log.Fatal("path to file is not provided")
	}
	var decoder compression.Decoder

	method := cmd.Flag("method").Value.String()

	switch method {
	case "vlc":
		decoder = vlc.New()
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

	compressed := decoder.Decode(data)

	err = os.WriteFile(uncompressedFileName(filePath), []byte(compressed), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func uncompressedFileName(path string) string {
	fName := filepath.Base(path)
	return strings.TrimSuffix(fName, filepath.Ext(fName)) + "." + uncompressedExtension
}

func init() {
	cmd.AddCommand(uncompressCmd)

	uncompressCmd.Flags().StringP("method", "m", "", "uncompression method: vlc, etc")
	if err := uncompressCmd.MarkFlagRequired("method"); err != nil {
		panic("flag m or 'method' not provided " + err.Error())
	}
}
