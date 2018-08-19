package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/labstack/gommon/color"
)

func getClass(path string, name string) {
	url := "http://localhost:8080/predict"
	path, _ = filepath.Abs(path)
	var jsonStr = []byte(`{"service":"imageserv","parameters":{"input":{"width":224,"height":224},"output":{"best":1},"mllib":{"gpu":false}},"data":["` + path + `"]}`)
	// DEBUG
	//fmt.Println("Request: " + string(jsonStr))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Close = true
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		stopDeepDetect(name)
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	parsedResponse := parseResponse(body)
	color.Println(color.Yellow("[") + color.Cyan(filepath.Base(path)) + color.Yellow("]") + color.Yellow(" Response: ") + color.Green(parsedResponse))
	renameFile(path, name, parsedResponse)
}
