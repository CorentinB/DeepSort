package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	filetype "gopkg.in/h2non/filetype.v1"
)

func runRecursively(path string, name string) ([]string, error) {
	searchDir := path

	fileList := make([]string, 0)
	e := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return err
	})

	if e != nil {
		stopDeepDetect(name)
		panic(e)
	}

	for _, file := range fileList {
		buf, _ := ioutil.ReadFile(file)
		if filetype.IsImage(buf) {
			getClass(file, name)
		}
	}

	return fileList, nil
}
