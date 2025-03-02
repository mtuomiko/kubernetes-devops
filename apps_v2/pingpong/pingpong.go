package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CountStruct struct {
	Count int `json:"count"`
}

type Config struct {
	Port       string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
}

func main() {
	// Create a root context with cancellation
	rootCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := Config{
		Port:       getEnvOrDefault("PORT", "4000"),
		DbHost:     getEnvOrDefault("DB_HOST", "localhost"),
		DbPort:     getEnvOrDefault("DB_PORT", "5432"),
		DbUser:     getEnvOrDefault("DB_USER", "postgres"),
		DbPassword: getEnvOrDefault("DB_PASSWORD", "Hunter2"),
		DbName:     getEnvOrDefault("DB_NAME", "pingpong"),
	}
	cfgPrint, _ := json.MarshalIndent(cfg, "", "\t")
	fmt.Println("Config: " + string(cfgPrint)) // yes it shows the password

	// "postgres://<user>:<password>@<host>:<port>/<dbname>";
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName,
	)

	db, err := connectDbPool(rootCtx, connString)
	if err != nil {
		log.Panic(err)
	}
	initDb(rootCtx, db)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handleRequests(w, r, rootCtx, db)
	})

	log.Printf("Pingpong server starting in port %s", cfg.Port)
	http.ListenAndServe(":"+cfg.Port, nil)
}

func connectDbPool(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, err
	}

	err = dbpool.Ping(ctx)
	if err != nil {
		return nil, err
	}

	log.Println("DB pool connected")

	return dbpool, nil
}

func initDb(ctx context.Context, db *pgxpool.Pool) {
	log.Println("Checking DB state")
	// simple test to check if DB is in a state where we can read necessary data
	selectStatement := "SELECT count FROM pingpong.pings WHERE id = 1;"

	var count int
	row := db.QueryRow(ctx, selectStatement)
	err := row.Scan(&count)

	if err != nil {
		log.Println(err)
	} else {
		return // if no errors, it's fine
	}

	log.Println("Initializing schema")
	schemaStatement := "CREATE SCHEMA pingpong"
	_, err = db.Exec(ctx, schemaStatement)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Initializing pings table")
	createTableStatement := `
	CREATE TABLE pingpong.pings (
		id integer PRIMARY KEY,
		count integer
	);`
	_, err = db.Exec(ctx, createTableStatement)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Initializing pings count")
	insertRecordStatement := "INSERT INTO pingpong.pings (id, count) VALUES (1, 0);"
	_, err = db.Exec(ctx, insertRecordStatement)
	if err != nil {
		log.Fatal(err)
	}
}

func handleRequests(w http.ResponseWriter, r *http.Request, ctx context.Context, db *pgxpool.Pool) {
	log.Printf("%s: %s %s", time.Now().Format(time.RFC3339), r.Method, r.URL.Path)

	if r.Method != "GET" {
		http.NotFound(w, r)
		return
	}

	if r.URL.Path == "/ping" {
		count := getCount(ctx, db)
		io.WriteString(w, "pong "+strconv.Itoa(count)+"\n")
		count++
		saveCount(ctx, db, count)
	} else if r.URL.Path == "/count" {
		w.Header().Set("Content-Type", "application/json")
		count := getCount(ctx, db)
		countStruct := CountStruct{
			Count: count,
		}
		json.NewEncoder(w).Encode(countStruct)
	} else {
		http.NotFound(w, r)
	}
}

func getCount(ctx context.Context, db *pgxpool.Pool) int {
	log.Println("Reading count from DB")

	selectStatement := "SELECT count FROM pingpong.pings WHERE id = 1;"

	var count int
	row := db.QueryRow(ctx, selectStatement)
	if err := row.Scan(&count); err != nil {
		log.Println(err)
		return -1
	}
	return count
}

func saveCount(ctx context.Context, db *pgxpool.Pool, count int) {
	log.Printf("Updating count to %d", count)

	updateStatement := "UPDATE pingpong.pings SET count = $1 WHERE id = 1;"

	_, err := db.Exec(ctx, updateStatement, count)
	if err != nil {
		log.Println(err) // ¯\_(ツ)_/¯
	}
}

func getEnvOrDefault(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
