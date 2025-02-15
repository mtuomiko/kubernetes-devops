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
	r.SetTrustedProxies(nil)
	r.Static("/", "./public")

	log.Printf("Server starting in port %s", port)
	r.Run(":" + port)
}
