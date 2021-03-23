package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Port fallback
	port := "5678"
	if portEnv, ok := os.LookupEnv("PORT"); ok {
		port = portEnv
	}

	r := gin.Default()
	r.Static("/", "./public")

	log.Printf("Server started on port %s", port)
	r.Run(":" + port)
}
