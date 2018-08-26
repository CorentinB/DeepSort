package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"encoding/base64"
	"strings"
	"github.com/savaki/jq"
	"sync"
)

// Runs the image through the classification engine and
// returns a slice of tags
func (c *ClassificationService) Classify(path string, content []byte, wg *sync.WaitGroup) []string {
	defer wg.Done()
	url := arguments.URL + "/predict"

	// Send image over as base64
	dataStr := base64.StdEncoding.EncodeToString(content)

	req, err := http.NewRequest("POST", url, strings.NewReader(`{
		"service": "` + c.Id + `",
		"parameters": {
			"input": {
				"width":224,
				"height":224
			},
			"output": { "best":1 },
			"mllib": { "gpu":false }
		},
		"data":["` + dataStr + `"]
	}`))

	resp, err := httpClient.Do(req)
	if err != nil {
		logError("Unable to classify this file.", "["+filepath.Base(path)+"]")
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	tagsJson, _ := jq.Parse(".body.predictions.[0].classes.[0].cat")
	value, _ := tagsJson.Apply(body)
	class := strings.Split(string(value), " ")
	class = class[1:]

	return class
}
