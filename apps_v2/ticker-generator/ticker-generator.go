package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

type StatusStruct struct {
	UUID      string `json:"uuid"`
	Timestamp string `json:"timestamp"`
}

func main() {
	statusFilePath, ok := os.LookupEnv("STATUS_FILE")
	if !ok {
		log.Fatal("STATUS_FILE env var not found")
	}
	log.Printf("Using status file path: %s", statusFilePath)

	staticUuid, err := uuid.NewUUID()
	if err != nil {
		os.Exit(1)
	}

	status := getStatus(staticUuid)
	interval := 5 * time.Second
	ticker := time.NewTicker(interval)
	// initial save
	saveStatus(status, statusFilePath)

	for range ticker.C {
		status = getStatus(staticUuid)
		log.Printf("Updating status: %s", status)

		saveStatus(status, statusFilePath)
	}
}

func getStatus(uuid uuid.UUID) StatusStruct {
	return StatusStruct{
		UUID:      uuid.String(),
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

func saveStatus(status StatusStruct, path string) {
	statusJson, err := json.Marshal(status)
	if err != nil {
		log.Println(err)
		return
	}
	err = os.WriteFile(path, statusJson, os.ModePerm)
	if err != nil {
		log.Println(err)
	}
}
