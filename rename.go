package main

import (
	"strings"
	"path/filepath"
	"crypto/md5"
	"encoding/hex"
	"github.com/labstack/gommon/color"
	"os"
	"bytes"
	"io"
)

var replaceSpace = strings.NewReplacer(" ", "_")
var replaceComma = strings.NewReplacer(",", "")
var replaceDoubleQuote = strings.NewReplacer("\"", "")

// FormatFileName takes an old path of an image,
// its content as bytes and its tags (e.g. from Classify)
// and returns a new file name for an image along with
// the formatted tag portion of the new name
func FormatFileName(oldName string, image []byte, tags []string) (newName string, tagPart string) {
	tagPart = formatTags(tags)

	hashBytes := md5.Sum(image)
	hash := hex.EncodeToString(hashBytes[:])
	extension := oldName[len(oldName)-4:]
	newName = tagPart + "_" + hash + extension

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

// commitFile copies or renames the target image
// to its new path
func commitFile(c *ClassificationService,
	path string, image []byte, tags []string) {

	// Generate file name
	newName, tagPart := FormatFileName(path, image, tags)

	// Log file name to console
	if len(newName) > 19 {
		// File name is too long, truncate it
		var message bytes.Buffer

		// Write first and last path of the file name
		truncatedName := newName[0:5] + "…" + newName[len(newName)-9:]
		message.WriteString(color.Yellow("[") + color.Cyan(truncatedName) + color.Yellow("]"))

		// Write tags
		message.WriteString(color.Yellow(" Response: ") + color.Green(tagPart))

		logSuccess(message.String(), c.Tag)
	} else {
		// File name fits in the console
		var message bytes.Buffer

		// Write file name
		message.WriteString(color.Yellow("[") + color.Cyan(newName) + color.Yellow("]"))

		// Pad to 19 characters
		for i := 15 - len(newName); i > 0; i-- {
			message.WriteByte(' ')
		}

		message.WriteString(color.Yellow(" Response: ") + color.Green(tagPart))

		// Write tag
		logSuccess(message.String(), c.Tag)
	}

	// Don't do anything if it's a dry run
	if arguments.DryRun { return }

	// Copy or rename the file
	if arguments.Output != "" {
		copyFile(path, newName, image)
	} else {
		renameFile(path, newName)
	}
}

// copyFile copies the image to a new path
func copyFile(oldPath string, newName string, image []byte) {
	// Get relative path of image to input path
	relPath, err := filepath.Rel(arguments.Input, oldPath)
	if err != nil { panic(err) }

	// Get relative directory of the image
	relDir := filepath.Dir(relPath)

	// Get new target directory of the image
	newDir := filepath.Join(arguments.Output, relDir)

	// Create directory if it doesn't exist
	err = os.MkdirAll(newDir, 0755)
	if err != nil {
		logError("Unable to create target directory this file.", "["+relPath+"]")
		os.Exit(1)
	}

	// Create new image file
	newPath := filepath.Join(newDir, newName)

	newImage, err := os.OpenFile(newPath,
		os.O_CREATE | os.O_EXCL | os.O_WRONLY, 0644)

	if err != nil {
		logError("Unable to open this file.", "["+newPath+"]")
		os.Exit(1)
	}

	defer newImage.Close()

	// Copy
	_, err = io.Copy(newImage, bytes.NewReader(image))

	if err != nil {
		logError("Failed to copy this file.", "["+newPath+"]")
		os.Exit(1)
	}
}

// renameFile renames the image
func renameFile(oldPath string, newImage string) {
	dir := filepath.Dir(oldPath)
	newPath := filepath.Join(dir, newImage)

	err := os.Rename(oldPath, newPath)
	if err != nil {
		logError("Unable to rename this file.", "["+filepath.Base(oldPath)+"]")
		os.Exit(1)
	}
}
