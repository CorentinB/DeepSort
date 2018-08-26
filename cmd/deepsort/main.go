package main

import (
	"os"
	"time"
	"github.com/labstack/gommon/color"
	"net/http"
	"log"
	"path/filepath"
)

var arguments struct {
	Input     string
	URL       string
	DryRun    bool
	Recursive bool
	Jobs      int
	Network   string
	CountDone int32
}

var httpClient = http.Client{}

func main() {
	start := time.Now()
	arguments.CountDone = 0
	arguments.Jobs = 1
	argumentParsing(os.Args)

	var service *ClassificationService

	switch arguments.Network {
	case "resnet-50":
		service = &resNet50
	case "googlenet":
		service = &googleNet
	default:
		panic("invalid service")
	}

	service.start()
	if arguments.Recursive {
		if arguments.Jobs == 1 {
			logSuccess("Starting image classification recursively..", service.Tag)
		} else {
			logSuccess("Starting image classification recursively with " +
				color.Green(arguments.Jobs) +
				color.Yellow(" parallel jobs.."), service.Tag)
		}

		// Open input folder
		f, err := os.Open(arguments.Input)
		if err != nil { log.Fatal(err) }
		defer f.Close()
		fileList, err := f.Readdirnames(-1)
		if err != nil {
			logError("Unable to process this directory.", "["+arguments.Input+"]")
			os.Exit(1)
		}

		service.process(fileList)
	} else {
		if arguments.Jobs == 1 {
			logSuccess("Starting image classification..", service.Tag)
		} else {
			logSuccess("Starting image classification with " +
				color.Green(arguments.Jobs) +
				color.Yellow(" parallel jobs.."), service.Tag)
		}

		// Open input folder
		fileList := make([]string, 0)
		e := filepath.Walk(arguments.Input, func(path string, f os.FileInfo, err error) error {
			fileList = append(fileList, path)
			return err
		})
		if e != nil {
			logError("Unable to process this directory.", "["+arguments.Input+"]")
			os.Exit(1)
		}

		service.process(fileList)
	}

	logSuccess(color.Yellow(arguments.CountDone)+
		color.Cyan(" pictures sorted in ")+
		color.Yellow(time.Since(start))+
		color.Cyan("!"), color.Cyan("Congrats,"))
}
