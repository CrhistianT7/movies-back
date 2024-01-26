package main

import (
	"fmt"
	"net/http"
)

func (app *application) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Hello")
}
