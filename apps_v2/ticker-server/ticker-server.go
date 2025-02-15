package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type StatusStruct struct {
	UUID      string `json:"uuid"`
	Timestamp string `json:"timestamp"`
}

func main() {
	port := "4000"
	if portEnv, ok := os.LookupEnv("PORT"); ok {
		port = portEnv
	}

	statusFilePath, ok := os.LookupEnv("STATUS_FILE")
	if !ok {
		log.Fatal("STATUS_FILE env var not found")
	}
	log.Printf("Using status file path: %s", statusFilePath)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handleStatus(w, r, statusFilePath)
	})

	log.Printf("Status server started in port %s", port)
	http.ListenAndServe(":"+port, nil)
}

func handleStatus(w http.ResponseWriter, r *http.Request, statusFilePath string) {
	log.Printf("%s %s", r.Method, r.URL.Path)
	if r.Method == "GET" && r.URL.Path == "/status" {
		statusJson, err := json.Marshal(readStatus(statusFilePath))
		if err != nil {
			fmt.Println(err)
		}
		w.Write(statusJson)
	} else {
		http.NotFound(w, r)
	}
}

func readStatus(path string) StatusStruct {
	jsonData, err := os.ReadFile(path)
	if err != nil {
		log.Println(err)
	}

	var status StatusStruct
	json.Unmarshal(jsonData, &status)
	return status
}
