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
	port := "4000"
	if portEnv, ok := os.LookupEnv("PORT"); ok {
		port = portEnv
	}

	randomString, err := uuid.NewUUID()

	if err != nil {
		os.Exit(1)
	}

	status := ""
	interval := 5 * time.Second
	go updateAndLogStatus(randomString, &status, interval)

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		handleStatus(w, r, &status)
	})

	log.Printf("Status server started in port %s", port)
	http.ListenAndServe(":"+port, nil)
}

func updateAndLogStatus(randomString uuid.UUID, status *string, interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		out := fmt.Sprintf("%s: %s", time.Now().Format(time.RFC3339), randomString)
		// update status variable back so status request can access the current value
		*status = out
		log.Println(out)
	}
}

func handleStatus(w http.ResponseWriter, r *http.Request, s *string) {
	if r.Method == "GET" {
		w.Write([]byte(*s))
	}
}
