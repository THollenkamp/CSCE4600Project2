package builtins

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Type(w io.Writer, args ...string) error {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	if len(args) != 1 {
		fmt.Println("Please provide a file name as an argument")
		return nil
	}

	fileName := args[0]

	filePath := filepath.Join(cwd, fileName)

	extension := strings.Replace(filepath.Ext(filePath), ".", "", 1)

	fmt.Printf("File type of %s: %s\n", fileName, extension)
	return nil
}
