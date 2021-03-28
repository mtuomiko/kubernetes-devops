package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Config struct {
	SharedDir     string
	PublicDir     string
	ImagePath     string
	ImageDataPath string
	ImageLink     string
	FetchURL      string
}

type Image struct {
	FetchTime time.Time `json:"fetchTime"`
}

func main() {
	sharedDir := filepath.Join(".", "shared")
	publicDir := filepath.Join(".", "public")
	config := Config{
		SharedDir:     sharedDir,
		PublicDir:     publicDir,
		ImagePath:     filepath.Join(sharedDir, "image.jpg"),
		ImageDataPath: filepath.Join(sharedDir, "image.json"),
		FetchURL:      "https://picsum.photos/400",
	}
	// Port fallback
	port := "5678"
	if portEnv, ok := os.LookupEnv("PORT"); ok {
		port = portEnv
	}

	fetchTime := readImageFetchTime(&config)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(CheckImageStatus(&fetchTime, &config))
	e.File("/image.jpg", config.ImagePath)
	e.Static("/", config.PublicDir)

	log.Printf("Server started on port %s", port)
	e.Logger.Fatal(e.Start(":" + port))
}

// Echo middleware which checks if the image needs to be refreshed
func CheckImageStatus(fetchTime *time.Time, config *Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			now := time.Now()
			if fetchTime.IsZero() || !equalDate(*fetchTime, now) {
				saveImage(config)
				saveImageData(config, now)
				*fetchTime = now
			}
			return next(c)
		}
	}
}

func readImageFetchTime(config *Config) time.Time {
	_, err := os.Open(config.ImagePath)
	if err != nil {
		return time.Time{}
	}

	imageDataJson, err := os.ReadFile(config.ImageDataPath)
	if err != nil {
		return time.Time{}
	}

	var imageData Image
	json.Unmarshal(imageDataJson, &imageData)

	return imageData.FetchTime
}

func equalDate(time1, time2 time.Time) bool {
	y1, m1, d1 := time1.Date()
	y2, m2, d2 := time2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

func saveImage(config *Config) {
	res, err := http.Get(config.FetchURL)
	if err != nil {
		log.Println(err)
		return
	}
	if res.StatusCode <= 199 || res.StatusCode >= 300 {
		log.Println("Image fetch failed. Resulted in status: " + res.Status)
		return
	}
	defer res.Body.Close()

	file, err := os.Create(config.ImagePath)
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

func saveImageData(config *Config, time time.Time) {
	imageData := Image{
		FetchTime: time,
	}
	imageDataJson, err := json.Marshal(imageData)
	if err != nil {
		log.Println(err)
		return
	}
	if err = ioutil.WriteFile(config.ImageDataPath, imageDataJson, os.ModePerm); err != nil {
		log.Println(err)
	}
}
