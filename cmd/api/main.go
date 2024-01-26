package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type application struct {
	Domain string
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

	// connect to database
	app.Domain = "example.com"
	log.Println("Starting server on port:", port)

	// start a live server
	server := &http.Server{
		Handler: app.routes(),
		Addr:    ":" + port,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Something when wrong when starting the server:", err)
	}
}
