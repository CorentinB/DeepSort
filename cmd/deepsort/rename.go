package main

import (
	"os"
	"path/filepath"
	"strings"
	"github.com/labstack/gommon/color"
	"crypto/md5"
	"encoding/hex"
	"sync/atomic"
)

var replaceSpace = strings.NewReplacer(" ", "_")
var replaceComma = strings.NewReplacer(",", "")
var replaceDoubleQuote = strings.NewReplacer("\"", "")

func (c *ClassificationService) Rename(path string, content []byte, tags []string) {
	response := formatTags(tags)

	if len(filepath.Base(path)) > 17 {
		name := filepath.Base(path)
		truncatedName := name[0:5] + "..." + name[len(name)-9:]
		logSuccess(color.Yellow("[")+color.Cyan(truncatedName)+
			color.Yellow("]")+color.Yellow(" Response: ")+
			color.Green(response), c.Tag)
	} else {
		logSuccess(color.Yellow("[")+color.Cyan(filepath.Base(path))+
			color.Yellow("]")+color.Yellow(" Response: ")+
			color.Green(response), c.Tag)
	}
	if arguments.DryRun != true {
		hashBytes := md5.Sum(content)
		hash := hex.EncodeToString(hashBytes[:])
		renameFile(path, hash, response)
	}

	// Atomic increment to prevent race conditions
	atomic.AddInt32(&arguments.CountDone, 1)
}

// Formats tags to a nice file name
func formatTags(class []string) string {
	result := strings.Join(class, " ")
	result = replaceSpace.Replace(result)
	result = replaceComma.Replace(result)
	result = replaceDoubleQuote.Replace(result)
	return result
}

func renameFile(path string, hash string, response string) {
	absPath, _ := filepath.Abs(path)
	dirPath := filepath.Dir(absPath)
	extension := path[len(path)-4:]
	newName := response + "_" + hash + extension
	err := os.Rename(absPath, dirPath+"/"+newName)
	if err != nil {
		logError("Unable to rename this file.", "["+filepath.Base(path)+"]")
		os.Exit(1)
	}
}
