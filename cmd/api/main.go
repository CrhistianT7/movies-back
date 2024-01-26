package main

import (
	"fmt"
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
	log.Println("Starting application on port:", port)
	// start a live server
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		log.Fatal("Something when wrong starting the server:", err)
	}
}
