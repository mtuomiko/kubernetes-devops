package main

import (
	"encoding/json"
	"fmt"
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
	greeting := ""
	if greetingEnv, ok := os.LookupEnv("GREETING"); ok {
		greeting = greetingEnv
	}

	dirPath := filepath.Join(".", "shared")
	statusPath := filepath.Join(dirPath, "status.json")
	countUrl := "http://pingpong-svc.exercises:5500/pingpongs"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handleStatus(w, r, statusPath, countUrl, greeting)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		handleHealthCheck(w, r, countUrl)
	})

	log.Printf("Status server started in port %s", port)
	http.ListenAndServe(":"+port, nil)
}

func handleStatus(
	w http.ResponseWriter,
	r *http.Request,
	statusPath string,
	countUrl string,
	greeting string,
) {
	if r.Method == "GET" {
		status := readStatus(statusPath)
		count, err := readCount(countUrl)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte(greeting + "\n" + status + "\n" + "Ping / Pongs: " + strconv.Itoa(count) + "\n"))
		}
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

func readCount(path string) (int, error) {
	res, err := http.Get(path)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	if res.StatusCode <= 199 || res.StatusCode >= 400 {
		log.Println("Get failed. Resulted in status: " + res.Status)
		return 0, fmt.Errorf("get failed, resulted in status: %d ", res.StatusCode)
	}

	var count Count
	if err := json.NewDecoder(res.Body).Decode(&count); err != nil {
		log.Println(err)
		return 0, err
	}
	return count.Count, nil
}

func handleHealthCheck(w http.ResponseWriter, r *http.Request, countUrl string) {
	if r.Method == "GET" {
		_, err := readCount(countUrl)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte("OK"))
		}
	} else {
		http.NotFound(w, r)
	}
}
