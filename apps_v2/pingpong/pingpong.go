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
	port := "4000"
	if portEnv, ok := os.LookupEnv("PORT"); ok {
		port = portEnv
	}
	countFilePath, ok := os.LookupEnv("COUNT_FILE")
	if !ok {
		log.Fatal("COUNT_FILE env var not found")
	}
	log.Printf("Using count file path: %s", countFilePath)

	counter := getOrInitCount(countFilePath)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlePing(w, r, &counter, countFilePath)
	})

	log.Printf("Pingpong server started in port %s", port)
	http.ListenAndServe(":"+port, nil)
}

func handlePing(w http.ResponseWriter, r *http.Request, counter *int, countFilePath string) {
	logMsg := fmt.Sprintf("%s: %s %s", time.Now().Format(time.RFC3339), r.Method, r.URL.Path)
	log.Println(logMsg)

	if r.Method == "GET" && r.URL.Path == "/ping" {
		io.WriteString(w, "pong "+strconv.Itoa(*counter)+"\n")
		*counter++
		saveCount(*counter, countFilePath)
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
