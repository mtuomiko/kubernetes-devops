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
	Count     int    `json:"pingPongCount"`
}

type TickerStruct struct {
	UUID      string `json:"uuid"`
	Timestamp string `json:"timestamp"`
}

type CountStruct struct {
	Count int `json:"count"`
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

	countFilePath, ok := os.LookupEnv("COUNT_FILE")
	if !ok {
		log.Fatal("COUNT env var not found")
	}
	log.Printf("Using count file path: %s", countFilePath)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handleStatus(w, r, statusFilePath, countFilePath)
	})

	log.Printf("Status server starting in port %s", port)
	http.ListenAndServe(":"+port, nil)
}

func handleStatus(w http.ResponseWriter, r *http.Request, statusFilePath string, countFilePath string) {
	log.Printf("%s %s", r.Method, r.URL.Path)
	if r.Method == "GET" && r.URL.Path == "/status" {
		ticker := readTicker(statusFilePath)
		count := readCount(countFilePath)

		status := &StatusStruct{
			UUID:      ticker.UUID,
			Timestamp: ticker.Timestamp,
			Count:     count.Count,
		}

		statusJson, err := json.Marshal(status)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "500 infernal server error", http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(statusJson)
	} else {
		http.NotFound(w, r)
	}
}

func readTicker(path string) TickerStruct {
	jsonData, err := os.ReadFile(path)
	if err != nil {
		log.Println(err)
	}

	var ticker TickerStruct
	json.Unmarshal(jsonData, &ticker)
	return ticker
}

func readCount(path string) CountStruct {
	log.Printf("starting to read count from %s", path)
	jsonData, err := os.ReadFile(path)
	if err != nil {
		log.Println(err)
	}

	var count CountStruct
	json.Unmarshal(jsonData, &count)
	return count
}
