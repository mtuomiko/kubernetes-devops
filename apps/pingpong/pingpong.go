package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// Port fallback
	port := "5500"
	if portEnv, ok := os.LookupEnv("PORT"); ok {
		port = portEnv
	}

	counter := 0

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlePing(w, r, &counter)
	})

	log.Printf("Pingpong server started in port %s", port)
	http.ListenAndServe(":"+port, nil)
}

func handlePing(w http.ResponseWriter, r *http.Request, c *int) {
	if r.Method == "GET" && r.URL.Path == "/" {
		io.WriteString(w, "pong "+strconv.Itoa(*c)+"\n")
		*c++
	} else {
		http.NotFound(w, r)
	}

}
