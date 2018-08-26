package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"crypto/md5"
	"encoding/base64"
	"encoding/hex"

	"github.com/CorentinB/DeepSort/pkg/logging"
	"github.com/labstack/gommon/color"
)

func googleNetClassification(path string, content []byte, arguments *Arguments, wg *sync.WaitGroup) {
	defer wg.Done()
	url := arguments.URL + "/predict"
	dataStr := base64.StdEncoding.EncodeToString(content)
	var jsonStr = []byte(`{"service":"deepsort-resnet","parameters":{"input":{"width":224,"height":224},"output":{"best":1},"mllib":{"gpu":false}},"data":["` + dataStr + `"]}`)
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
		hashBytes := md5.Sum(content)
		hash := hex.EncodeToString(hashBytes[:])
		renameFile(path, hash, arguments, parsedResponse)
	}
	arguments.CountDone++
}

func resNet50Classification(path string, content []byte, arguments *Arguments, wg *sync.WaitGroup) {
	defer wg.Done()
	url := arguments.URL + "/predict"
	dataStr := base64.StdEncoding.EncodeToString(content)
	var jsonStr = []byte(`{"service":"deepsort-resnet-50","parameters":{"input":{"width":224,"height":224},"output":{"best":1},"mllib":{"gpu":false}},"data":["` + dataStr + `"]}`)
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
		hashBytes := md5.Sum(content)
		hash := hex.EncodeToString(hashBytes[:])
		renameFile(path, hash, arguments, parsedResponse)
	}
	arguments.CountDone++
}
