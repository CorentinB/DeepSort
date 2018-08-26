package main

import (
	"io/ioutil"
	"sync"
	"gopkg.in/h2non/filetype.v1"
)

func (c *ClassificationService) process(fileList []string) {
	// Synchronization
	count := 0
	var wg sync.WaitGroup

	// Process files in the input folder
	for _, file := range fileList {
		buf, _ := ioutil.ReadFile(file)
		if filetype.IsImage(buf) {
			count++
			wg.Add(1)

			tags := c.Classify(file, buf, &wg)
			c.Rename(file, buf, tags)

			if count == arguments.Jobs {
				wg.Wait()
				count = 0
			}
		}
	}
}
