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

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Config struct {
	ImagePath     string
	ImageMetaPath string
	FetchURL      string
}

type Image struct {
	FetchTime time.Time `json:"fetchTime"`
}

type Todo struct {
	ID    string `json:"id" validate:"required"`
	Title string `json:"title" validate:"required"`
}

var todos = map[string]*Todo{
	"88422cbc-90b0-11eb-a8b3-0242ac130003": {
		ID:    "88422cbc-90b0-11eb-a8b3-0242ac130003",
		Title: "Get a haircut",
	},
	"8d284fd6-90b0-11eb-a8b3-0242ac130003": {
		ID:    "8d284fd6-90b0-11eb-a8b3-0242ac130003",
		Title: "Get a real job",
	},
}

var validate *validator.Validate

func main() {
	config := Config{
		ImagePath:     filepath.Join(".", "data", "image.jpg"),
		ImageMetaPath: filepath.Join(".", "data", "image.json"),
		FetchURL:      "https://picsum.photos/400",
	}
	// Port fallback
	port := "5600"
	if portEnv, ok := os.LookupEnv("PORT"); ok {
		port = portEnv
	}

	fetchTime := readImageFetchTime(&config)
	validate = validator.New()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.File("/image.jpg", config.ImagePath, CheckImageStatus(&fetchTime, &config))
	e.GET("/todos", getAllTodos)
	e.POST("/todos", createTodo)

	log.Printf("Server started on port %s", port)
	e.Logger.Fatal(e.Start(":" + port))
}

func readImageFetchTime(config *Config) time.Time {
	_, err := os.Open(config.ImagePath)
	if err != nil {
		return time.Time{}
	}

	imageDataJson, err := os.ReadFile(config.ImageMetaPath)
	if err != nil {
		return time.Time{}
	}

	var imageData Image
	json.Unmarshal(imageDataJson, &imageData)

	return imageData.FetchTime
}

// Echo middleware which checks if the image needs to be refreshed
func CheckImageStatus(fetchTime *time.Time, config *Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			now := time.Now()
			if fetchTime.IsZero() || !equalDate(*fetchTime, now) {
				saveImage(config)
				saveImageMeta(config, now)
				*fetchTime = now
			}
			return next(c)
		}
	}
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
	if res.StatusCode <= 199 || res.StatusCode >= 400 {
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

func saveImageMeta(config *Config, time time.Time) {
	imageMeta := Image{
		FetchTime: time,
	}
	imageMetaJson, err := json.Marshal(imageMeta)
	if err != nil {
		log.Println(err)
		return
	}
	if err = ioutil.WriteFile(config.ImageMetaPath, imageMetaJson, os.ModePerm); err != nil {
		log.Println(err)
	}
}

// Routes
func getAllTodos(c echo.Context) error {
	values := []*Todo{}
	for _, value := range todos {
		values = append(values, value)
	}
	return c.JSON(http.StatusOK, values)
}

func createTodo(c echo.Context) error {
	newTodo := &Todo{}
	if err := c.Bind(newTodo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if newTodo.ID != "" {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	newTodo.ID = uuid.NewString()
	if err := validate.Struct(newTodo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	todos[newTodo.ID] = newTodo
	return c.JSON(http.StatusCreated, newTodo)
}
