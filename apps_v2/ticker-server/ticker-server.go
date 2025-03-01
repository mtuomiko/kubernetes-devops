package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type StatusStruct struct {
	UUID        string `json:"uuid"`
	Timestamp   string `json:"timestamp"`
	Count       int    `json:"pingPongCount"`
	Message     string `json:"message"`
	Information string `json:"information"`
}

type TickerStruct struct {
	UUID      string `json:"uuid"`
	Timestamp string `json:"timestamp"`
}

type CountStruct struct {
	Count int `json:"count"`
}

type Config struct {
	StatusFilePath      string
	InformationFilePath string
	CountUri            string
	Message             string
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func main() {
	port := getEnvOrDefault("PORT", "4000")

	cfg := Config{
		StatusFilePath:      getEnvOrDefault("STATUS_FILE", "./status/status.json"),
		InformationFilePath: getEnvOrDefault("INFORMATION_FILE", "./config/information.txt"),
		CountUri:            getEnvOrDefault("COUNT_URI", "http://localhost:4001/count"),
		Message:             getEnvOrDefault("MESSAGE", "n/a"),
	}
	cfgPrint, _ := json.MarshalIndent(cfg, "", "\t")
	fmt.Println("Config: " + string(cfgPrint))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handleStatus(w, r, cfg)
	})

	log.Printf("ticker-server starting in port %s", port)
	http.ListenAndServe(":"+port, nil)
}

func handleStatus(w http.ResponseWriter, r *http.Request, cfg Config) {
	log.Printf("%s %s", r.Method, r.URL.Path)
	if r.Method == "GET" && r.URL.Path == "/status" {
		ticker := readTicker(cfg)
		information := readInformation(cfg)
		count := getCount(cfg)

		status := &StatusStruct{
			UUID:        ticker.UUID,
			Timestamp:   ticker.Timestamp,
			Count:       count,
			Message:     cfg.Message,
			Information: information,
		}

		statusJson, err := json.Marshal(status)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "500 infernal server error", http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(statusJson)
	} else {
		http.NotFound(w, r)
	}
}

func readTicker(cfg Config) TickerStruct {
	log.Printf("Reading %s", cfg.StatusFilePath)
	jsonData, err := os.ReadFile(cfg.StatusFilePath)
	if err != nil {
		log.Println(err)
	}

	var ticker TickerStruct
	json.Unmarshal(jsonData, &ticker)
	return ticker
}

func readInformation(cfg Config) string {
	information, err := os.ReadFile(cfg.InformationFilePath)
	if err != nil {
		log.Println(err)
		return "n/a"
	}

	return string(information)
}

func getCount(cfg Config) int {
	count := new(CountStruct)
	log.Printf("Calling %s", cfg.CountUri)
	getJson(cfg.CountUri, count)
	return count.Count
}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func getEnvOrDefault(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
