package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type application struct {
	Domain string
	DSN    string
	DB     *sql.DB
}

func main() {
	// set application config

	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in the enviroment")
	}

	var app application

	// read from command line
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5", "Postgres connectrin string")
	flag.Parse()

	// connect to database
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	app.DB = conn

	app.Domain = "example.com"
	log.Println("Starting server on port:", port)

	// start a live server
	server := &http.Server{
		Handler: app.routes(),
		Addr:    ":" + port,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("Something when wrong when starting the server:", err)
	}
}
