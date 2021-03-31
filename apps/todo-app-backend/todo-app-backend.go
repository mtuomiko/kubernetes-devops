package main

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Config struct {
	ImagePath     string
	ImageMetaPath string
	FetchURL      string
	DBHost        string
	DBPort        string // Gets used as string, no need for int conversion
	DBUser        string
	DBPassword    string
	DBName        string
}

type Image struct {
	FetchTime time.Time `json:"fetchTime"`
}

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title" validate:"required"`
}

var (
	cfg          Config
	validate     *validator.Validate
	ctx          = context.Background()
	pgdb         *pg.DB
	defaultTodos = []string{
		"Get a haircut",
		"Get a real job",
	}
)

func main() {
	cfg = Config{
		ImagePath:     filepath.Join(".", "data", "image.jpg"),
		ImageMetaPath: filepath.Join(".", "data", "image.json"),
		FetchURL:      "https://picsum.photos/400",
		DBHost:        getEnvOrDefault("POSTGRES_HOST", "localhost"),
		DBPort:        getEnvOrDefault("POSTGRES_PORT", "5432"),
		DBUser:        getEnvOrDefault("POSTGRES_USER", "postgres"),
		DBPassword:    os.Getenv("POSTGRES_PASSWORD"),
		DBName:        getEnvOrDefault("POSTGRES_DB", "postgres"),
	}
	// Port fallback
	port := getEnvOrDefault("PORT", "5600")

	fetchTime := readImageFetchTime()
	validate = validator.New()

	// Initialize DB connection
	pgdb = pg.Connect(&pg.Options{
		Addr:     cfg.DBHost + ":" + cfg.DBPort,
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
		Database: cfg.DBName,
	})
	defer pgdb.Close()
	if err := pgdb.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	// Verify schema and add table if needed
	if err := initSchema(); err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.File("/image.jpg", cfg.ImagePath, checkImageStatus(&fetchTime))
	e.GET("/todos", allTodosHandler)
	e.POST("/todos", createTodoHandler)

	log.Printf("Server started on port %s", port)
	e.Logger.Fatal(e.Start(":" + port))
}

func readImageFetchTime() time.Time {
	log.Println(cfg)
	_, err := os.Open(cfg.ImagePath)
	if err != nil {
		return time.Time{}
	}

	imageDataJson, err := os.ReadFile(cfg.ImageMetaPath)
	if err != nil {
		return time.Time{}
	}

	var imageData Image
	json.Unmarshal(imageDataJson, &imageData)

	return imageData.FetchTime
}

// Echo middleware which checks if the image needs to be refreshed
func checkImageStatus(fetchTime *time.Time) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			now := time.Now()
			if fetchTime.IsZero() || !equalDate(*fetchTime, now) {
				saveImage()
				saveImageMeta(now)
				*fetchTime = now
			}
			return next(c)
		}
	}
}

func saveImage() {
	res, err := http.Get(cfg.FetchURL)
	if err != nil {
		log.Println(err)
		return
	}
	if res.StatusCode <= 199 || res.StatusCode >= 400 {
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
	imageMeta := Image{
		FetchTime: time,
	}
	imageMetaJson, err := json.Marshal(imageMeta)
	if err != nil {
		log.Println(err)
		return
	}
	if err = ioutil.WriteFile(cfg.ImageMetaPath, imageMetaJson, os.ModePerm); err != nil {
		log.Println(err)
	}
}

// Route handlers
func allTodosHandler(c echo.Context) error {
	todos, err := getAllTodos()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, todos)
}

func createTodoHandler(c echo.Context) error {
	newTodo := &Todo{}
	if err := c.Bind(newTodo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if newTodo.ID != 0 {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := validate.Struct(newTodo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Everything should now be OK with the payload, failure after this is a DB problem?
	if err := insertTodo(newTodo); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusCreated, newTodo)
}

// DB stuff
func initSchema() error {
	// Select from non existing table should fail
	_, err := getAllTodos()
	if err != nil {
		if err = createSchema(); err != nil {
			return err
		}
		// Add default todos to an empty table
		for _, todoTitle := range defaultTodos {
			newTodo := &Todo{Title: todoTitle}
			if err = insertTodo(newTodo); err != nil {
				return err
			}
		}
	}
	return nil
}

func createSchema() error {
	err := pgdb.Model((*Todo)(nil)).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})
	if err != nil {
		return err
	}
	return nil
}

func getAllTodos() ([]Todo, error) {
	var todos []Todo
	if err := pgdb.Model(&todos).Select(); err != nil {
		return nil, err
	}
	return todos, nil
}

func insertTodo(todo *Todo) error {
	//newTodo := &Todo{Title: title}
	_, err := pgdb.Model(todo).Returning("id").Insert()
	if err != nil {
		return err
	}
	return nil
}

func getEnvOrDefault(envKey string, defaultStr string) string {
	if env, ok := os.LookupEnv(envKey); ok {
		return env
	}
	return defaultStr
}

func equalDate(time1, time2 time.Time) bool {
	y1, m1, d1 := time1.Date()
	y2, m2, d2 := time2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
