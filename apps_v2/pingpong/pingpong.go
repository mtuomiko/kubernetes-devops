package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type CountStruct struct {
	Count int `json:"count"`
}

func main() {
	port := getEnvOrDefault("PORT", "4000")

	countFilePath := getEnvOrDefault("COUNT_FILE", "./count/count.json")
	log.Printf("Using count file path: %s", countFilePath)

	counter := getOrInitCount(countFilePath)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handleRequests(w, r, &counter, countFilePath)
	})

	log.Printf("Pingpong server starting in port %s", port)
	http.ListenAndServe(":"+port, nil)
}

func handleRequests(w http.ResponseWriter, r *http.Request, counter *int, countFilePath string) {
	logMsg := fmt.Sprintf("%s: %s %s", time.Now().Format(time.RFC3339), r.Method, r.URL.Path)
	log.Println(logMsg)

	if r.Method != "GET" {
		http.NotFound(w, r)
		return
	}

	if r.URL.Path == "/ping" {
		io.WriteString(w, "pong "+strconv.Itoa(*counter)+"\n")
		*counter++
		saveCount(*counter, countFilePath)
	} else if r.URL.Path == "/count" {
		w.Header().Set("Content-Type", "application/json")
		countStruct := CountStruct{
			Count: *counter,
		}
		json.NewEncoder(w).Encode(countStruct)
	} else {
		http.NotFound(w, r)
	}
}

func getOrInitCount(countFilePath string) int {
	jsonData, err := os.ReadFile(countFilePath)

	if err == nil {
		var jsonCount CountStruct
		json.Unmarshal(jsonData, &jsonCount)
		return jsonCount.Count
	} else {
		// assume any errors are because file does not yet exist
		saveCount(0, countFilePath)
		return 0
	}
}

func saveCount(count int, countFilePath string) {
	countStruct := CountStruct{
		Count: count,
	}
	countJson, err := json.Marshal(countStruct)
	if err != nil {
		log.Println(err)
		return
	}
	err = os.WriteFile(countFilePath, countJson, os.ModePerm)
	if err != nil {
		log.Println(err)
	}
}

func getEnvOrDefault(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
