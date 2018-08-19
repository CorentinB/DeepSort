package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/CorentinB/DeepSort/pkg/logging"
	filetype "gopkg.in/h2non/filetype.v1"
)

func run(arguments *Arguments) {
	// Handle parallelization
	count := 0
	var wg sync.WaitGroup
	// Open input folder
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

	// Process files in the input folder
	for _, file := range files {
		path := arguments.Input + "/" + file.Name()
		buf, _ := ioutil.ReadFile(path)
		if filetype.IsImage(buf) {
			count++
			wg.Add(1)
			if arguments.Network == "resnet-50" {
				go resNet50Classification(path, arguments, &wg)
			} else {
				go googleNetClassification(path, arguments, &wg)
			}
			if count == arguments.Jobs {
				wg.Wait()
				count = 0
			}
		}
	}
}

func runRecursively(arguments *Arguments) ([]string, error) {
	// Handle parallelization
	count := 0
	var wg sync.WaitGroup
	// Open input folder
	fileList := make([]string, 0)
	e := filepath.Walk(arguments.Input, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return err
	})
	if e != nil {
		logging.Error("Unable to process this directory.", "["+arguments.Input+"]")
		os.Exit(1)
	}

	// Process files in the input folder
	for _, file := range fileList {
		buf, _ := ioutil.ReadFile(file)
		if filetype.IsImage(buf) {
			count++
			wg.Add(1)
			if arguments.Network == "resnet-50" {
				go resNet50Classification(file, arguments, &wg)
			} else {
				go googleNetClassification(file, arguments, &wg)
			}
		}
		if count == arguments.Jobs {
			wg.Wait()
			count = 0
		}
	}
	return fileList, nil
}
