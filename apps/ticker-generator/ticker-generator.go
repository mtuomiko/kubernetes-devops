package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

type StatusStruct struct {
	Status string `json:"status"`
}

func main() {
	randomString, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}

	dirPath := filepath.Join(".", "shared")
	path := filepath.Join(dirPath, "status.json")

	ticker := time.NewTicker(5 * time.Second)
	for range ticker.C {
		statusString := fmt.Sprintf("%s: %s", time.Now().Format(time.RFC3339), randomString)
		log.Println(statusString)
		saveStatus(statusString, path)
	}
}

func saveStatus(statusString string, path string) {
	status := StatusStruct{
		Status: statusString,
	}
	statusJson, err := json.Marshal(status)
	if err != nil {
		log.Println(err)
		return
	}
	err = ioutil.WriteFile(path, statusJson, os.ModePerm)
	if err != nil {
		log.Println(err)
	}
}
