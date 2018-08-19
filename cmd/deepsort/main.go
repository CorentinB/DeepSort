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
}

func main() {
	start := time.Now()
	arguments := new(Arguments)
	argumentParsing(os.Args, arguments)
	startGoogleNet(arguments)
	if arguments.Recursive == true {
		logging.Success("Starting image classification recursively..", "[GoogleNet]")
		runRecursively(arguments)
	} else {
		logging.Success("Starting image classification..", "[GoogleNet]")
		run(arguments)
	}
	color.Println(color.Cyan("Done in ") + color.Yellow(time.Since(start)) + color.Cyan("!"))
}
