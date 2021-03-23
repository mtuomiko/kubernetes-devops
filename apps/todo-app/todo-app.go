package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}

func main() {
	// Port fallback
	port := "5678"
	if portEnv, ok := os.LookupEnv("PORT"); ok {
		port = portEnv
	}

	http.HandleFunc("/", rootHandler)

	log.Printf("Server started in port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
