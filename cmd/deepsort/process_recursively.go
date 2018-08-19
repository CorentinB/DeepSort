package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/CorentinB/DeepSort/pkg/logging"
	filetype "gopkg.in/h2non/filetype.v1"
)

func runRecursively(arguments *Arguments) ([]string, error) {
	searchDir := arguments.Input

	fileList := make([]string, 0)
	e := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return err
	})

	if e != nil {
		logging.Error("Unable to process this directory.", "["+searchDir+"]")
		os.Exit(1)
	}

	for _, file := range fileList {
		buf, _ := ioutil.ReadFile(file)
		if filetype.IsImage(buf) {
			getClass(file, arguments)
		}
	}

	return fileList, nil
}
