package main

import (
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
)

func main() {
	randomString, err := uuid.NewUUID()

	if err != nil {
		os.Exit(1)
	}

	ticker := time.NewTicker(5 * time.Second)

	for range ticker.C {
		out := fmt.Sprintf("%s: %s", time.Now().Format(time.RFC3339), randomString)
		fmt.Println(out)
	}
}
