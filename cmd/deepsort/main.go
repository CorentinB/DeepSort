package main

import (
	"os"
	"time"

	"github.com/labstack/gommon/color"
)

func main() {
	start := time.Now()
	name := startDeepDetect(os.Args[1])
	color.Println(color.Yellow("[") + color.Cyan("CONTAINER: "+name) + color.Yellow("] ") + color.Yellow("Starting image classification.. "))
	runRecursively(os.Args[1], name)
	stopDeepDetect(name)
	color.Println(color.Yellow("[") + color.Cyan("CONTAINER: "+name) + color.Yellow("] ") + color.Green("Successfully stopped DeepDetect. "))
	color.Println(color.Cyan("Done in ") + color.Yellow(time.Since(start)) + color.Cyan("!"))
}
