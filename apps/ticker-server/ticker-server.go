package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type StatusStruct struct {
	Status string `json:"status"`
}

func main() {
	// Port fallback
	port := "4000"
	if portEnv, ok := os.LookupEnv("PORT"); ok {
		port = portEnv
	}

	dirPath := filepath.Join(".", "shared")
	path := filepath.Join(dirPath, "status.json")

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		handleStatus(w, r, path)
	})

	log.Printf("Status server started in port %s", port)
	http.ListenAndServe(":"+port, nil)
}

func handleStatus(w http.ResponseWriter, r *http.Request, path string) {
	if r.Method == "GET" {
		jsonData, err := os.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}
		var status StatusStruct
		json.Unmarshal(jsonData, &status)
		w.Write([]byte(status.Status + "\n"))
	} else {
		http.NotFound(w, r)
	}
}
