package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

type Count struct {
	Count int `json:"count"`
}

var (
	host        = "pingpong-db-svc.exercises"
	port        = 5432
	user        = "postgres"
	password    = os.Getenv("POSTGRES_PASSWORD")
	dbname      = "postgres"
	routePrefix = "/"
	jsonRoute   = routePrefix + "pingpongs"
	db          *sql.DB
	counter     int
)

func main() {
	psqlConfig := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	// Port fallback
	port := "5500"
	if portEnv, ok := os.LookupEnv("PORT"); ok {
		port = portEnv
	}

	// Start connecting to db
	go connectLoop(psqlConfig)
	defer db.Close()

	// Health check response
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if db == nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else if err := db.Ping(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			fmt.Fprintf(w, "OK")
		}
	})

	http.HandleFunc(routePrefix, func(w http.ResponseWriter, r *http.Request) {
		if db == nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			handlePing(w, r)
		}
	})
	http.HandleFunc(jsonRoute, func(w http.ResponseWriter, r *http.Request) {
		if db == nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			handleCount(w, r)
		}
	})

	log.Printf("Pingpong server started in port %s", port)
	http.ListenAndServe(":"+port, nil)
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" && r.URL.Path == routePrefix {
		io.WriteString(w, "pong "+strconv.Itoa(counter)+"\n")
		counter++
		updateCount(db, counter)
	} else {
		http.NotFound(w, r)
	}
}

func handleCount(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" && r.URL.Path == jsonRoute {
		countStruct := Count{
			Count: counter,
		}
		countJson, err := json.Marshal(countStruct)
		if err != nil {
			log.Println(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(countJson)
	} else {
		http.NotFound(w, r)
	}
}

func initTable(db *sql.DB) {
	log.Println("Initializing pings table")
	createTable := `
	CREATE TABLE pings (
		id SERIAL PRIMARY KEY,
		count INTEGER
	);`
	insertRecord := "INSERT INTO pings (count) VALUES (0);"
	_, err := db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(insertRecord)
	if err != nil {
		log.Fatal(err)
	}
}

func readCount(db *sql.DB) int {
	log.Println("Reading count from DB")
	selectStatement := "SELECT count FROM pings WHERE id = 1;"
	var count int
	row := db.QueryRow(selectStatement)
	if err := row.Scan(&count); err != nil {
		log.Println(err)
		initTable(db)
		return 0
	}
	return count
}

func updateCount(db *sql.DB, count int) {
	updateStatement := "UPDATE pings SET count = $1 WHERE id = 1;"
	_, err := db.Exec(updateStatement, count)
	if err != nil {
		log.Fatal(err)
	}
}

// https://alex.dzyoba.com/blog/go-connect-loop/
// keep trying to connect to db
func connectLoop(psqlConfig string) error {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {

		case <-ticker.C:
			log.Println("connecting to db...")
			dbConnection, err := sql.Open("postgres", psqlConfig)
			if err != nil {
				log.Println(err)
				continue
			}

			if err = dbConnection.Ping(); err != nil {
				log.Println(err)
				continue
			}

			db = dbConnection
			countVal := readCount(db)
			counter = countVal
			return nil
		}
	}
}
