package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/CorentinB/DeepSort/pkg/logging"
	"github.com/labstack/gommon/color"
)

func googleNetClassification(path string, arguments *Arguments, wg *sync.WaitGroup) {
	defer wg.Done()
	url := arguments.URL + "/predict"
	path, _ = filepath.Abs(path)
	var jsonStr = []byte(`{"service":"deepsort-resnet","parameters":{"input":{"width":224,"height":224},"output":{"best":1},"mllib":{"gpu":false}},"data":["` + path + `"]}`)
	// DEBUG
	//fmt.Println("Request: " + string(jsonStr))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Close = true
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logging.Error("Unable to classify this file.", "["+filepath.Base(path)+"]")
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	parsedResponse := parseResponse(body)
	if len(filepath.Base(path)) > 17 {
		name := filepath.Base(path)
		truncatedName := name[0:5] + "..." + name[len(name)-9:]
		logging.Success(color.Yellow("[")+color.Cyan(truncatedName)+
			color.Yellow("]")+color.Yellow(" Response: ")+
			color.Green(parsedResponse), "[GoogleNet]")
	} else {
		logging.Success(color.Yellow("[")+color.Cyan(filepath.Base(path))+
			color.Yellow("]")+color.Yellow(" Response: ")+
			color.Green(parsedResponse), "[GoogleNet]")
	}
	if arguments.DryRun != true {
		if isValid(arguments.Output) == false {
			renameFile(path, arguments, parsedResponse)
		} else {
			fmt.Println("Output selected")
		}
	}
	arguments.CountDone++
}

func resNet50Classification(path string, arguments *Arguments, wg *sync.WaitGroup) {
	defer wg.Done()
	url := arguments.URL + "/predict"
	path, _ = filepath.Abs(path)
	var jsonStr = []byte(`{"service":"deepsort-resnet-50","parameters":{"input":{"width":224,"height":224},"output":{"best":1},"mllib":{"gpu":false}},"data":["` + path + `"]}`)
	// DEBUG
	//fmt.Println("Request: " + string(jsonStr))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Close = true
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logging.Error("Unable to classify this file.", "["+filepath.Base(path)+"]")
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	parsedResponse := parseResponse(body)
	if len(filepath.Base(path)) > 17 {
		name := filepath.Base(path)
		truncatedName := name[0:5] + "..." + name[len(name)-9:]
		logging.Success(color.Yellow("[")+color.Cyan(truncatedName)+
			color.Yellow("]")+color.Yellow(" Response: ")+
			color.Green(parsedResponse), "[ResNet-50]")
	} else {
		logging.Success(color.Yellow("[")+color.Cyan(filepath.Base(path))+
			color.Yellow("]")+color.Yellow(" Response: ")+
			color.Green(parsedResponse), "[ResNet-50]")
	}
	if arguments.DryRun != true {
		if isValid(arguments.Output) == false &&
			arguments.OutputChoice == false {
			renameFile(path, arguments, parsedResponse)
		} else if arguments.OutputChoice == true &&
			isValid(arguments.Output) == false {
			logging.Error("Wrong output folder.", "")
			os.Exit(1)
		}
	}
	arguments.CountDone++
}
