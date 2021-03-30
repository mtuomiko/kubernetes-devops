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
	countUrl := "http://pingpong-svc.exercises:5501/pingpongs"

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		handleStatus(w, r, statusPath, countUrl)
	})

	log.Printf("Status server started in port %s", port)
	http.ListenAndServe(":"+port, nil)
}

func handleStatus(
	w http.ResponseWriter,
	r *http.Request,
	statusPath string,
	countUrl string,
) {
	if r.Method == "GET" {
		status := readStatus(statusPath)
		count := readCount(countUrl)
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
	res, err := http.Get(path)
	if err != nil {
		log.Println(err)
		return 0
	}
	if res.StatusCode <= 199 || res.StatusCode >= 400 {
		log.Println("Get failed. Resulted in status: " + res.Status)
		return 0
	}

	var count Count
	if err := json.NewDecoder(res.Body).Decode(&count); err != nil {
		log.Println(err)
		return 0
	}
	return count.Count
}
