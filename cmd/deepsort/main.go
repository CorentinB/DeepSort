package main

import (
	"os"
	"time"

	"github.com/CorentinB/DeepSort/pkg/logging"
	"github.com/labstack/gommon/color"
)

type Arguments struct {
	Input  string
	URL    string
	DryRun bool
}

func main() {
	start := time.Now()
	arguments := new(Arguments)
	argumentParsing(os.Args, arguments)
	startGoogleNet(arguments)
	logging.Success("Starting image classification..", "[GoogleNet]")
	runRecursively(arguments)
	color.Println(color.Cyan("Done in ") + color.Yellow(time.Since(start)) + color.Cyan("!"))
}
