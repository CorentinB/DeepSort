package main

import (
	"os"
	"time"

	"github.com/CorentinB/DeepSort/pkg/logging"
	"github.com/labstack/gommon/color"
)

type Arguments struct {
	Input     string
	URL       string
	DryRun    bool
	Recursive bool
	Jobs      int
}

func main() {
	start := time.Now()
	arguments := new(Arguments)
	arguments.Jobs = 1
	argumentParsing(os.Args, arguments)
	startGoogleNet(arguments)
	if arguments.Recursive == true {
		if arguments.Jobs == 1 {
			logging.Success("Starting image classification recursively..", "[GoogleNet]")
		} else {
			logging.Success("Starting image classification recursively with "+string(arguments.Jobs)+" parallel jobs..", "[GoogleNet]")
		}
		runRecursively(arguments)
	} else {
		if arguments.Jobs == 1 {
			logging.Success("Starting image classification..", "[GoogleNet]")
		} else {
			logging.Success("Starting image classification with "+string(arguments.Jobs)+" parallel jobs..", "[GoogleNet]")
		}
		run(arguments)
	}
	color.Println(color.Cyan("Done in ") + color.Yellow(time.Since(start)) + color.Cyan("!"))
}
