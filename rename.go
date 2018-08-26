package main

import (
	"strings"
	"path/filepath"
	"crypto/md5"
	"encoding/hex"
	"github.com/labstack/gommon/color"
	"os"
	"bytes"
)

var replaceSpace = strings.NewReplacer(" ", "_")
var replaceComma = strings.NewReplacer(",", "")
var replaceDoubleQuote = strings.NewReplacer("\"", "")

// FormatFileName takes an old path of an image,
// its content as bytes and its tags (e.g. from Classify)
// and returns a new file name for an image along with
// the formatted tag portion of the new name
func FormatFileName(path string, image []byte, tags []string) (fullPath string, tagPart string) {
	tagPart = formatTags(tags)

	hashBytes := md5.Sum(image)
	hash := hex.EncodeToString(hashBytes[:])
	absPath, _ := filepath.Abs(path)
	dirPath := filepath.Dir(absPath)
	extension := path[len(path)-4:]
	newName := tagPart + "_" + hash + extension
	fullPath = filepath.Join(dirPath, newName)

	return
}

// Formats tags to be used in a file name
func formatTags(class []string) string {
	result := strings.Join(class, " ")
	result = replaceSpace.Replace(result)
	result = replaceComma.Replace(result)
	result = replaceDoubleQuote.Replace(result)
	return result
}

func renameFile(c *ClassificationService,
	path string, image []byte, tags []string) {

	// Generate file name
	newPath, tagPart := FormatFileName(path, image, tags)

	name := filepath.Base(path)

	// Log file name to console
	if len(name) > 19 {
		// File name is too long, truncate it
		var message bytes.Buffer

		// Write first and last path of the file name
		truncatedName := name[0:5] + "…" + name[len(name)-9:]
		message.WriteString(color.Yellow("[") + color.Cyan(truncatedName) + color.Yellow("]"))

		// Write tags
		message.WriteString(color.Yellow(" Response: ") + color.Green(tagPart))

		logSuccess(message.String(), c.Tag)
	} else {
		// File name fits in the console
		var message bytes.Buffer

		// Write file name
		message.WriteString(color.Yellow("[") + color.Cyan(name) + color.Yellow("]"))

		// Pad to 19 characters
		for i := 15 - len(name); i > 0; i-- {
			message.WriteByte(' ')
		}

		message.WriteString(color.Yellow(" Response: ") + color.Green(tagPart))

		// Write tag
		logSuccess(message.String(), c.Tag)
	}

	// Rename file
	if !arguments.DryRun {
		err := os.Rename(path, newPath)
		if err != nil {
			logError("Unable to rename this file.", "["+filepath.Base(path)+"]")
			os.Exit(1)
		}
	}
}
