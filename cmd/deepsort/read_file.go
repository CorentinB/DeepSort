package main

import (
	"bytes"
	"encoding/base64"
	"github.com/CorentinB/DeepSort/pkg/logging"
	"io"
	"os"
	"path/filepath"
)

// Reads file and returns data string for DeepDetect
func readFile(filePath string) string {
	f, err := os.Open(filePath)
	if err != nil {
		logging.Error("Can't open the file.", "["+filepath.Base(filePath)+"]")
		os.Exit(1)
	}
	defer f.Close()

	var buf bytes.Buffer
	enc := base64.NewEncoder(base64.StdEncoding, &buf)

	if _, err := io.Copy(enc, f); err != nil {
		logging.Error("Can't read the file.", "["+filepath.Base(filePath)+"]")
		os.Exit(1)
	}

	enc.Close()
	return buf.String()
}
