package command

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/skhe55/archiver-go/internal/services"
	"github.com/spf13/cobra"
)

var vlcCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Compress file using variable-length code",
	Run:   compress,
}

const compressedExtension = "vlc"

func compress(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		log.Fatal("path to file is not provided")
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

	compressed := services.Encode(string(data))

	err = os.WriteFile(compressedFileName(filePath), []byte(compressed), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func compressedFileName(path string) string {
	fName := filepath.Base(path)
	return strings.TrimSuffix(fName, filepath.Ext(fName)) + "." + compressedExtension
}

func init() {
	compressCmd.AddCommand(vlcCmd)
}
