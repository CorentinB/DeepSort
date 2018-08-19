package main

import (
	"bytes"
	"net/http"
	"os"

	"github.com/CorentinB/DeepSort/pkg/logging"
)

func startGoogleNet(arguments *Arguments) {
	// Starting the image classification service
	logging.Success("Starting the classification service..", "[GoogleNet]")
	url := arguments.URL + "/services/deepsort-googlenet"
	var jsonStr = []byte(`{"mllib":"caffe","description":"DeepSort-GoogleNet","type":"supervised","parameters":{"input":{"connector":"image"},"mllib":{"nclasses":1000}},"model":{"repository":"/opt/models/ggnet/"}}`)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logging.Error("Error while starting the classification service, please check if DeepDetect is running.", "[GoogleNet]")
		os.Exit(1)
	}
	defer resp.Body.Close()
	if resp.Status != "201 Created" && resp.Status != "500 Internal Server Error" {
		logging.Error("Error while starting the classification service, please check if DeepDetect is running.", "[GoogleNet]")
		os.Exit(1)
	}
	if resp.Status == "500 Internal Server Error" {
		logging.Success("Looks like you already have the deepsort-googlenet service started, no need to create a new one.", "[GoogleNet]")
	} else {
		logging.Success("Successfully started the image classification service.", "[GoogleNet]")
	}
}
