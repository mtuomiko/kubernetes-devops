package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

func main() {
	// Port fallback
	port := "4000"
	if portEnv, ok := os.LookupEnv("PORT"); ok {
		port = portEnv
	}

	randomString, err := uuid.NewUUID()

	if err != nil {
		os.Exit(1)
	}

	status := ""
	go logStatus(randomString, &status)

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		handleStatus(w, r, &status)
	})

	log.Printf("Status server started in port %s", port)
	http.ListenAndServe(":"+port, nil)
}

func handleStatus(w http.ResponseWriter, r *http.Request, s *string) {
	if r.Method == "GET" {
		w.Write([]byte(*s))
	}
}

func logStatus(randomString uuid.UUID, s *string) {
	ticker := time.NewTicker(5 * time.Second)

	for range ticker.C {
		out := fmt.Sprintf("%s: %s", time.Now().Format(time.RFC3339), randomString)
		*s = out
		log.Println(out)
	}
}
