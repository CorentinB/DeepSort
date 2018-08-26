package main

import (
	"bytes"
	"net/http"
	"os"
)

type ClassificationService struct{
	Conn *http.Client
	Id  string
	Tag string
	Description string
	Repository  string
}

var googleNet = ClassificationService{
	Id:  "deepsort-googlenet",
	Tag: "[GoogleNet]",
	Description: "DeepSort-GoogleNet",
	Repository:  "/opt/models/ggnet/",
}

var resNet50 = ClassificationService{
	Id:  "deepsort-resnet-50",
	Tag: "[ResNet-50]",
	Description: "DeepSort-ResNet-50",
	Repository:  "/opt/models/resnet_50/",
}

func (c *ClassificationService) start() {
	// Starting the image classification service
	logSuccess("Starting the classification service..", c.Tag)
	url := arguments.URL + "/services/" + c.Id

	var jsonStr = []byte(`{
		"mllib": "caffe",
		"description": "` + c.Description + `",
		"type": "supervised",
		"parameters": {
			"input": { "connector": "image" },
			"mllib": { "nclasses": 1000 }
		},
		"model": { "repository": "` + c.Repository + `" }
	}`)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	resp, err := httpClient.Do(req)
	if err != nil {
		logError("Error while starting the classification service, " +
			"please check if DeepDetect is running.", c.Tag)
		os.Exit(1)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 201:
		logSuccess("Successfully started the" +
			"image classification service.", c.Tag)
	case 500:
		logSuccess("Looks like you already have the " + c.Id +
			" service started, no need to create a new one.", c.Tag)
	default:
		logError("Error while starting the classification service, " +
			"please check if DeepDetect is running.", c.Tag)
		os.Exit(1)
	}
}
