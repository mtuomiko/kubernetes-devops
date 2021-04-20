package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Config struct {
	PublicPort     string
	MetricsPort    string
	RandomImageURL string
}

var (
	cfg         Config
	catsCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "cat_view_counter",
		Help: "Number of cats viewed",
	})
	dogsCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "dog_view_counter",
		Help: "Number of dogs viewed",
	})
	randomCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "random_view_counter",
		Help: "Number of random images viewed",
	})
)

func init() {
	prometheus.MustRegister(catsCounter)
	prometheus.MustRegister(dogsCounter)
	prometheus.MustRegister(randomCounter)
}

func main() {
	cfg = Config{
		PublicPort:     getEnvOrDefault("PORT", "5700"),
		MetricsPort:    "8081",
		RandomImageURL: "https://picsum.photos/400",
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.File("/cat", "assets/cat.jpg", incrementCounter(catsCounter))
	e.File("/dog", "assets/dog.jpg", incrementCounter(dogsCounter))

	e.GET("/random", randomImageHandler)

	go func() {
		metrics := echo.New()
		metrics.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
		metrics.Logger.Fatal(metrics.Start(":8090"))
	}()

	log.Printf("Server started on port %s", cfg.PublicPort)
	e.Logger.Fatal(e.Start(":" + cfg.PublicPort))
}

// Some routes use File handler
func incrementCounter(counter prometheus.Counter) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			counter.Inc()
			return next(c)
		}
	}
}

func randomImageHandler(c echo.Context) error {
	randomCounter.Inc()
	res, err := http.Get(cfg.RandomImageURL)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	if res.StatusCode <= 199 || res.StatusCode >= 400 {
		log.Println("Image fetch failed. Resulted in status: " + res.Status)
		echo.NewHTTPError(http.StatusInternalServerError)
	}
	defer res.Body.Close()

	return c.Stream(http.StatusOK, "image/jpeg", res.Body)
}

func getEnvOrDefault(envKey string, defaultStr string) string {
	if env, ok := os.LookupEnv(envKey); ok {
		return env
	}
	return defaultStr
}
