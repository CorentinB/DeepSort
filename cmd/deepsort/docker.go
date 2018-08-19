package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/labstack/gommon/color"
)

func startDeepDetect(path string) string {
	// Starting the DeepDetect docker image
	name := randString(11, charset)
	path, _ = filepath.Abs(path)
	cmd := "docker"
	args := []string{"run", "-d", "-p", "8080:8080", "-v", path + ":" + path, "--name", name, "beniz/deepdetect_cpu"}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		stopDeepDetect(name)
		os.Exit(1)
	}
	color.Println(color.Yellow("[") + color.Cyan("CONTAINER: "+name) + color.Yellow("] ") + color.Green("Successfully started DeepDetect. "))
	// Starting the image classification service
	time.Sleep(time.Second)
	url := "http://localhost:8080/services/imageserv"
	var jsonStr = []byte(`{"mllib":"caffe","description":"image classification service","type":"supervised","parameters":{"input":{"connector":"image"},"mllib":{"nclasses":1000}},"model":{"repository":"/opt/models/ggnet/"}}`)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		stopDeepDetect(name)
		panic(err)
	}
	defer resp.Body.Close()
	if resp.Status != "201 Created" {
		stopDeepDetect(name)
		os.Exit(1)
	}
	color.Println(color.Yellow("[") + color.Cyan("CONTAINER: "+name) + color.Yellow("] ") + color.Green("Successfully started the image classification service. "))
	return name
}

func stopDeepDetect(name string) {
	// Stopping the DeepDetect docker image
	cmd := "docker"
	args := []string{"stop", name}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
