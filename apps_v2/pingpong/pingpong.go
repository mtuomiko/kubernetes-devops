package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	port := "4000"
	if portEnv, ok := os.LookupEnv("PORT"); ok {
		port = portEnv
	}

	counter := 0

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		handlePing(w, r, &counter)
	})

	log.Printf("Pingpong server started in port %s", port)
	http.ListenAndServe(":"+port, nil)
}

func handlePing(w http.ResponseWriter, r *http.Request, c *int) {
	logMsg := fmt.Sprintf("%s: %s %s", time.Now().Format(time.RFC3339), r.Method, r.URL.Path)
	log.Println(logMsg)

	if r.Method == "GET" && r.URL.Path == "/ping" {
		io.WriteString(w, "pong "+strconv.Itoa(*c)+"\n")
		*c++
	} else {
		http.NotFound(w, r)
	}
}
