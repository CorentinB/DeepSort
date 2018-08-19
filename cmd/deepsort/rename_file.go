package main

import (
	"os"
	"path/filepath"

	"github.com/CorentinB/DeepSort/pkg/logging"
)

func renameFile(path string, arguments *Arguments, response string) {
	absPath, _ := filepath.Abs(path)
	hash := hashFileMD5(absPath)
	dirPath := filepath.Dir(absPath)
	extension := path[len(path)-4:]
	newName := response + "_" + hash + extension

	err := os.Rename(absPath, dirPath+"/"+newName)
	if err != nil {
		logging.Error("Unable to rename this file.", "["+filepath.Base(path)+"]")
		os.Exit(1)
	}
}
