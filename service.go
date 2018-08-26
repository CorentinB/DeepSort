package DeepSort

import (
	"net/http"
	"bytes"
)

type ClassificationService struct{
	Conn *http.Client
	Url  string
	Id   string
	Tag  string
	Description string
}

// Connects to DeepDetect and loads a classification service
// from the specified repository.
func (c *ClassificationService) Load(repo string) error {
	// Starting the image classification service
	url := c.Url + "/services/" + c.Id

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
