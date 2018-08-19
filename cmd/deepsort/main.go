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
	Network   string
	CountDone int
}

func main() {
	start := time.Now()
	arguments := new(Arguments)
	arguments.CountDone = 0
	arguments.Jobs = 1
	argumentParsing(os.Args, arguments)
	if arguments.Network == "resnet-50" {
		startResNet50(arguments)
		if arguments.Recursive == true {
			if arguments.Jobs == 1 {
				logging.Success("Starting image classification recursively..", "[ResNet-50]")
			} else {
				logging.Success("Starting image classification recursively with "+
					color.Green(arguments.Jobs)+
					color.Yellow(" parallel jobs.."), "[ResNet-50]")
			}
			runRecursively(arguments)
		} else {
			if arguments.Jobs == 1 {
				logging.Success("Starting image classification..", "[ResNet-50]")
			} else {
				logging.Success("Starting image classification with "+
					color.Green(arguments.Jobs)+
					color.Yellow(" parallel jobs.."), "[ResNet-50]")
			}
			run(arguments)
		}
	} else {
		startGoogleNet(arguments)
		if arguments.Recursive == true {
			if arguments.Jobs == 1 {
				logging.Success("Starting image classification recursively..", "[GoogleNet]")
			} else {
				logging.Success("Starting image classification recursively with "+
					color.Green(arguments.Jobs)+
					color.Yellow(" parallel jobs.."), "[GoogleNet]")
			}
			runRecursively(arguments)
		} else {
			if arguments.Jobs == 1 {
				logging.Success("Starting image classification..", "[GoogleNet]")
			} else {
				logging.Success("Starting image classification with "+
					color.Green(arguments.Jobs)+
					color.Yellow(" parallel jobs.."), "[GoogleNet]")
			}
			run(arguments)
		}
	}
	logging.Success(color.Yellow(arguments.CountDone)+
		color.Cyan(" pictures sorted in ")+
		color.Yellow(time.Since(start))+
		color.Cyan("!"), color.Cyan("Congrats,"))
}
