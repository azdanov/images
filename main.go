package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"log"
	"mime"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("No arguments specified.")
	}

	image := os.Args[1]

	file, err := os.Open(image)
	if os.IsNotExist(err) {
		log.Fatal("File not found.")
	}

	mimeType := mime.TypeByExtension(filepath.Ext(os.Args[1]))

	supported := []string{"image/svg", "image/svg+xml", "image/gif", "image/jpeg", "image/png"}

	if !inSlice(mimeType, supported) {
		log.Fatal("Image type not supported.")
	}

	info, _ := file.Stat()
	buf := make([]byte, info.Size())

	fReader := bufio.NewReader(file)
	_, _ = fReader.Read(buf)

	data := fmt.Sprintf("data:%s;base64,%s", mimeType, base64.StdEncoding.EncodeToString(buf))

	fmt.Print(data)
}

func inSlice(a string, list []string) bool {
	for _, b := range list {
		if a == b {
			return true
		}
	}
	return false
}
