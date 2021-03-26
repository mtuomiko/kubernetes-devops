package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type Count struct {
	Count int `json:"count"`
}

func main() {
	// Port fallback
	port := "5500"
	if portEnv, ok := os.LookupEnv("PORT"); ok {
		port = portEnv
	}

	dirPath := filepath.Join(".", "shared")
	path := filepath.Join(dirPath, "count.json")

	counter := 0
	jsonData, err := os.ReadFile(path)
	if err == nil {
		var jsonCount Count
		json.Unmarshal(jsonData, &jsonCount)
		counter = jsonCount.Count
	} else {
		saveCount(counter, path)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlePing(w, r, &counter, path)
	})

	log.Printf("Pingpong server started in port %s", port)
	http.ListenAndServe(":"+port, nil)
}

func handlePing(w http.ResponseWriter, r *http.Request, count *int, path string) {
	if r.Method == "GET" && r.URL.Path == "/" {
		io.WriteString(w, "pong "+strconv.Itoa(*count)+"\n")
		*count++
		saveCount(*count, path)
	} else {
		http.NotFound(w, r)
	}
}

func saveCount(count int, path string) {
	countStruct := Count{
		Count: count,
	}
	countJson, err := json.Marshal(countStruct)
	if err != nil {
		log.Println(err)
		return
	}
	err = ioutil.WriteFile(path, countJson, os.ModePerm)
	if err != nil {
		log.Println(err)
	}
}
