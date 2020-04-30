package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func createDbConnectionString(config *config) string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", config.DatabaseUser, config.DatabasePassword, config.DatabaseHost, config.DatabasePort, config.DatabaseName)
}

func main() {
	log.Println("Starting application ...")

	config := newConfig()
	err := config.getFromEnv()
	if err != nil {
		log.Fatalf("Unable to get config from env: %v", err.Error())
	}

	config.prettyLog()

	connectionString := createDbConnectionString(config)
	log.Printf("Using the following connection string for: %v", connectionString)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("Unable to establish connection to the database: %v", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Unable to ping database :%v ", err.Error())
	}

	log.Println("Successfully connected to MySQL database!")

	// Static files:
	http.Handle("/", http.FileServer(http.Dir(config.StaticPath)))

	// Health Check
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		err = db.Ping()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("ko"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	log.Println("Listening on port 3000 ...")
	http.ListenAndServe(":3000", nil)
}
