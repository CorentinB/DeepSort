package DeepSort

import (
	"encoding/base64"
	"net/http"
	"strings"
	"io/ioutil"
	"github.com/savaki/jq"
)

// Runs the image through the classification engine and
// returns a slice of tags
func (c *ClassificationService) Classify(image []byte) ([]string, error) {
	url := c.Url + "/predict"

	// Send image over as base64
	dataStr := base64.StdEncoding.EncodeToString(image)

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

	resp, err := c.Conn.Do(req)
	if err != nil { return nil, err }
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	tagsJson, _ := jq.Parse(".body.predictions.[0].classes.[0].cat")
	value, _ := tagsJson.Apply(body)
	class := strings.Split(string(value), " ")
	class = class[1:]

	return class, nil
}
