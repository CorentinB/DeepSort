package main

import (
	"os"
	"path/filepath"
	"github.com/labstack/gommon/color"
	"github.com/CorentinB/DeepSort"
	"bytes"
)


func renameFile(c *DeepSort.ClassificationService,
	path string, image []byte, tags []string) {

	// Generate file name
	newPath, tagPart := DeepSort.FormatFileName(path, image, tags)

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
