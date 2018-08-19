package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/CorentinB/DeepSort/pkg/logging"
	filetype "gopkg.in/h2non/filetype.v1"
)

func run(arguments *Arguments) {
	f, err := os.Open(arguments.Input)
	if err != nil {
		log.Fatal(err)
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		logging.Error("Unable to process this directory.", "["+arguments.Input+"]")
		os.Exit(1)
	}

	for _, file := range files {
		path := arguments.Input + "/" + file.Name()
		buf, _ := ioutil.ReadFile(path)
		if filetype.IsImage(buf) {
			googleNetClassification(path, arguments)
		}
	}
}

func runRecursively(arguments *Arguments) ([]string, error) {
	fileList := make([]string, 0)
	e := filepath.Walk(arguments.Input, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return err
	})

	if e != nil {
		logging.Error("Unable to process this directory.", "["+arguments.Input+"]")
		os.Exit(1)
	}

	for _, file := range fileList {
		buf, _ := ioutil.ReadFile(file)
		if filetype.IsImage(buf) {
			googleNetClassification(file, arguments)
		}
	}

	return fileList, nil
}
