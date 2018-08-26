package main

import (
	"io/ioutil"
	"gopkg.in/h2non/filetype.v1"
	"path/filepath"
	"os"
	"github.com/CorentinB/DeepSort"
)

func process(c *DeepSort.ClassificationService, fileList []string) {
	// Process files in the input folder
	for _, file := range fileList {
		buf, _ := ioutil.ReadFile(file)
		if filetype.IsImage(buf) {
			tags, err := c.Classify(buf)
			if err != nil {
				tag := "["+filepath.Base(file)+"]"
				logError("Unable to classify this file.", tag)
				os.Exit(1)
			}

			renameFile(c, file, buf, tags)
		}
	}
}
