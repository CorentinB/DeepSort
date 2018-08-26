package main

import (
	"os"
	"github.com/CorentinB/DeepSort"
)

func startService(c *DeepSort.ClassificationService) {
	var repo string

	switch arguments.Network {
	case "resnet-50":
		c.ID = "deepsort-resnet-50"
		c.Tag = "[ResNet-50]"
		c.Description = "DeepSort-ResNet-50"
		repo = "/opt/models/resnet_50/"
	case "googlenet":
		c.ID = "deepsort-googlenet"
		c.Tag = "[GoogleNet]"
		c.Description = "DeepSort-GoogleNet"
		repo = "/opt/models/ggnet/"
	default:
		panic("invalid service")
	}

	logSuccess("Starting the classification service..", c.Tag)
	err := c.Load(repo)
	switch err {
	case nil:
		logSuccess("Successfully started the" +
			"image classification service.", c.Tag)

	case DeepSort.ErrAlreadyRunning:
		logSuccess("Looks like you already have the " + c.ID+
			" service started, no need to create a new one.", c.Tag)

	case DeepSort.ErrStartFailed: fallthrough
	default:
		logError("Error while starting the classification service, " +
			"please check if DeepDetect is running.", c.Tag)
		os.Exit(1)
	}
}
