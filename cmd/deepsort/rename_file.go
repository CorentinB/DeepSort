package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func renameFile(path string, name string, response string) {
	absPath, _ := filepath.Abs(path)
	hash := hashFileMD5(absPath, name)
	dirPath := filepath.Dir(absPath)
	extension := path[len(path)-4:]
	newName := response + "_" + hash + extension

	err := os.Rename(absPath, dirPath+"/"+newName)
	if err != nil {
		fmt.Println(err)
		stopDeepDetect(name)
		return
	}
}
