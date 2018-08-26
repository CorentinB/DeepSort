package main

import (
	"io/ioutil"
	"gopkg.in/h2non/filetype.v1"
	"path/filepath"
	"os"
)

func process(c *ClassificationService, fileList []string) (processed int) {
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

			commitFile(c, file, buf, tags)
			processed++
		}
	}

	return processed
}
