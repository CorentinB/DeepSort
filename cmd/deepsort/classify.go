package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/CorentinB/DeepSort/pkg/logging"
	"github.com/labstack/gommon/color"
)

func googleNetClassification(path string, arguments *Arguments) {
	url := arguments.URL + "/predict"
	path, _ = filepath.Abs(path)
	var jsonStr = []byte(`{"service":"deepsort-googlenet","parameters":{"input":{"width":224,"height":224},"output":{"best":1},"mllib":{"gpu":false}},"data":["` + path + `"]}`)
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
	color.Println(color.Yellow("[") + color.Cyan(filepath.Base(path)) + color.Yellow("]") + color.Yellow(" Response: ") + color.Green(parsedResponse))
	renameFile(path, arguments, parsedResponse)
}
