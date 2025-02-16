package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Interval      time.Duration
	FetchURL      string
	ImagePath     string
	ImageMetaPath string
}

type ImageMeta struct {
	FetchTime time.Time `json:"fetchTime"`
}

var (
	cfg       Config
	fetchTime time.Time
)

func main() {
	port := getEnvOrDefault("PORT", "5678")

	imageDir := getEnvOrDefault("IMAGE_DIR", "./images")
	cfg = Config{
		// go thinks int * Duration is not ok, fine
		Interval:      time.Duration(getEnvToIntOrDefault("INTERVAL_SECONDS", 3600)) * time.Second,
		ImagePath:     path.Join(imageDir, "image.jpg"),
		ImageMetaPath: path.Join(imageDir, "image.json"),
		FetchURL:      "https://picsum.photos/400",
	}

	log.Printf("Using refresh interval of %s", cfg.Interval)

	// Get state on startup. Afterwards fetchTime is in-memory only for checking if needs to be refreshed. The meta json
	// file is updated on image refresh.
	fetchTime = readImageFetchTime()

	if imageNeedsRefresh() {
		log.Printf("Initial refresh at start of app, old fetchTime %s", fetchTime)
		refreshImage()
	} else {
		log.Printf("No image refresh at start of app, fetchTime %s", fetchTime)
	}

	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.Use(static.Serve("/", static.LocalFile("./public", false)))
	r.GET("/image.jpg", CheckRefresh, func(c *gin.Context) {
		c.File(cfg.ImagePath)
	})

	log.Printf("Server starting in port %s", port)
	r.Run(":" + port)
}

func readImageFetchTime() time.Time {
	_, err := os.Open(cfg.ImagePath)
	if err != nil {
		return time.Time{}
	}

	imageMetaJson, err := os.ReadFile(cfg.ImageMetaPath)
	if err != nil {
		return time.Time{}
	}

	var imageMeta ImageMeta
	json.Unmarshal(imageMetaJson, &imageMeta)

	return imageMeta.FetchTime
}

func imageNeedsRefresh() bool {
	now := time.Now()
	cutoff := fetchTime.Add(cfg.Interval)
	result := cutoff.Before(now)
	if result {
		log.Printf("Image older than %s", cfg.Interval)
	}
	return result
}

func refreshImage() {
	getAndSaveImage()
	fetchTime = time.Now()
	saveImageMeta(fetchTime)
}

func getAndSaveImage() {
	log.Printf("Getting image from %s to %s", cfg.FetchURL, cfg.ImagePath)
	res, err := http.Get(cfg.FetchURL)
	if err != nil {
		log.Println(err)
		return
	}
	if res.StatusCode <= 199 || res.StatusCode >= 300 {
		log.Println("Image fetch failed. Resulted in status: " + res.Status)
		return
	}
	defer res.Body.Close()

	file, err := os.Create(cfg.ImagePath)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		log.Println(err)
	}
}

func saveImageMeta(time time.Time) {
	imageMeta := ImageMeta{
		FetchTime: time,
	}
	imageMetaJson, err := json.Marshal(imageMeta)
	if err != nil {
		log.Println(err)
		return
	}
	if err = os.WriteFile(cfg.ImageMetaPath, imageMetaJson, os.ModePerm); err != nil {
		log.Println(err)
	}
}

// gin middleware
func CheckRefresh(c *gin.Context) {
	if imageNeedsRefresh() {
		refreshImage()
	}

	c.Next()
}

func getEnvOrDefault(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvToIntOrDefault(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		if i, err := strconv.Atoi(value); err == nil {
			return i
		}
	}

	return fallback
}
