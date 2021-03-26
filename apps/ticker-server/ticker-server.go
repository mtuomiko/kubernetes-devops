package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type StatusStruct struct {
	Status string `json:"status"`
}

type Count struct {
	Count int `json:"count"`
}

func main() {
	// Port fallback
	port := "4000"
	if portEnv, ok := os.LookupEnv("PORT"); ok {
		port = portEnv
	}

	dirPath := filepath.Join(".", "shared")
	statusPath := filepath.Join(dirPath, "status.json")
	countPath := filepath.Join(dirPath, "count.json")

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		handleStatus(w, r, statusPath, countPath)
	})

	log.Printf("Status server started in port %s", port)
	http.ListenAndServe(":"+port, nil)
}

func handleStatus(
	w http.ResponseWriter,
	r *http.Request,
	statusPath string,
	countPath string,
) {
	if r.Method == "GET" {
		status := readStatus(statusPath)
		count := readCount(countPath)
		w.Write([]byte(status + "\n" + "Ping / Pongs: " + strconv.Itoa(count) + "\n"))
	} else {
		http.NotFound(w, r)
	}
}

func readStatus(path string) string {
	statusJson, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	var status StatusStruct
	json.Unmarshal(statusJson, &status)
	return status.Status
}

func readCount(path string) int {
	statusJson, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	var count Count
	json.Unmarshal(statusJson, &count)
	return count.Count
}
