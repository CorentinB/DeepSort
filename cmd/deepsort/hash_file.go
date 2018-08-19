package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"os"
)

func hashFileMD5(filePath string, name string) string {
	h := md5.New()
	f, err := os.Open(filePath)
	if err != nil {
		stopDeepDetect(name)
		log.Fatal(err)
	}
	defer f.Close()
	if _, err := io.Copy(h, f); err != nil {
		stopDeepDetect(name)
		log.Fatal(err)
	}
	return (hex.EncodeToString(h.Sum(nil)))
}
