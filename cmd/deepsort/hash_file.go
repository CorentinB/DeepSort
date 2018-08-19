package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"

	"github.com/CorentinB/DeepSort/pkg/logging"
)

func hashFileMD5(filePath string) string {
	h := md5.New()
	f, err := os.Open(filePath)
	if err != nil {
		logging.Error("Can't open the file.", "["+filepath.Base(filePath)+"]")
		os.Exit(1)
	}
	defer f.Close()
	if _, err := io.Copy(h, f); err != nil {
		logging.Error("Can't get checksum.", "["+filepath.Base(filePath)+"]")
		os.Exit(1)
	}
	return (hex.EncodeToString(h.Sum(nil)))
}
