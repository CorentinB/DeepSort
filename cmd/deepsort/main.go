package main

import (
	"os"
	"time"
	"github.com/labstack/gommon/color"
	"net/http"
	"log"
	"path/filepath"
	"github.com/CorentinB/DeepSort"
)

func main() {
	start := time.Now()
	argumentParsing(os.Args)

	// Start a new classification service
	var c = DeepSort.ClassificationService{
		Conn: &http.Client{},
		Url: arguments.URL,
	}
	startService(&c)

	var fileList []string

	if arguments.Recursive {
		if arguments.Jobs == 1 {
			logSuccess("Starting image classification recursively..", c.Tag)
		} else {
			logSuccess("Starting image classification recursively with " +
				color.Green(arguments.Jobs) +
				color.Yellow(" parallel jobs.."), c.Tag)
		}

		// Open input folder
		f, err := os.Open(arguments.Input)
		if err != nil { log.Fatal(err) }
		defer f.Close()
		fileList, err = f.Readdirnames(-1)
		if err != nil {
			logError("Unable to process this directory.", "["+arguments.Input+"]")
			os.Exit(1)
		}
	} else {
		if arguments.Jobs == 1 {
			logSuccess("Starting image classification..", c.Tag)
		} else {
			logSuccess("Starting image classification with " +
				color.Green(arguments.Jobs) +
				color.Yellow(" parallel jobs.."), c.Tag)
		}

		// Open input folder
		e := filepath.Walk(arguments.Input, func(path string, f os.FileInfo, err error) error {
			fileList = append(fileList, path)
			return err
		})
		if e != nil {
			logError("Unable to process this directory.", "["+arguments.Input+"]")
			os.Exit(1)
		}
	}

	process(&c, fileList)

	logSuccess(color.Yellow(len(fileList))+
		color.Cyan(" pictures sorted in ")+
		color.Yellow(time.Since(start))+
		color.Cyan("!"), color.Cyan("Congrats,"))
}
