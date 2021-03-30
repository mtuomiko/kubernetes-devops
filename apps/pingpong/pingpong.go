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

	_ "github.com/lib/pq"
)

type Count struct {
	Count int `json:"count"`
}

var (
	host     = "pingpong-db-svc.exercises"
	port     = 5432
	user     = "postgres"
	password = os.Getenv("POSTGRES_PASSWORD")
	dbname   = "postgres"
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

	// Init DB connection
	db, err := sql.Open("postgres", psqlConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("DB connection ok")
	counter := readCount(db)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlePing(w, r, &counter, db)
	})
	http.HandleFunc("/pingpongs", func(w http.ResponseWriter, r *http.Request) {
		handleCount(w, r, counter)
	})

	log.Printf("Pingpong server started in port %s", port)
	http.ListenAndServe(":"+port, nil)
}

func handlePing(w http.ResponseWriter, r *http.Request, count *int, db *sql.DB) {
	if r.Method == "GET" && r.URL.Path == "/" {
		io.WriteString(w, "pong "+strconv.Itoa(*count)+"\n")
		*count++
		updateCount(db, *count)
	} else {
		http.NotFound(w, r)
	}
}

func handleCount(w http.ResponseWriter, r *http.Request, count int) {
	if r.Method == "GET" && r.URL.Path == "/pingpongs" {
		countStruct := Count{
			Count: count,
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
