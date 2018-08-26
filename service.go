package DeepSort

import (
	"net/http"
	"bytes"
)

// ClassificationService describes a DeepDetect service
// that can be used to classify images.
// ClassificationService.Load() must be called before usage.
type ClassificationService struct{
	// The HTTP client used for all API calls
	Conn *http.Client
	// The base URL of the API
	URL  string
	// ID used for registering it in DeepDetect
	ID   string
	// Tag used for logging
	Tag  string
	// Description used for DeepDetect
	Description string
}

// Load connects to DeepDetect and loads a
// classification service from the specified repository
// using the specified values of the receiving object.
func (c *ClassificationService) Load(repo string) error {
	// Starting the image classification service
	url := c.URL + "/services/" + c.ID

	var jsonStr = []byte(`{
		"mllib": "caffe",
		"description": "` + c.Description + `",
		"type": "supervised",
		"parameters": {
			"input": { "connector": "image" },
			"mllib": { "nclasses": 1000 }
		},
		"model": { "repository": "` + repo + `" }
	}`)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	resp, err := c.Conn.Do(req)
	if err != nil { return ErrStartFailed }
	defer resp.Body.Close()

	switch resp.StatusCode {
		case 201: return nil
		case 500: return ErrAlreadyRunning
		default:  return ErrStartFailed
	}
}
