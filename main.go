package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// Initialize a template page with index.html using html/template package

var indexPath = "/app/index.html"
var tpl = template.Must(template.ParseFiles(indexPath))

// Create the handler for the root index page
func indexHandler(w http.ResponseWriter, r *http.Request) {

	tpl = template.Must(template.ParseFiles(indexPath))

	// This template is executed using html/template package
	tpl.Execute(w, nil)
}

func main() {

	// Load the .env file in current directory
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file.")
		os.Setenv("PORT", "3000")
	}

	// Get the port from the environment variable
	port := os.Getenv("PORT")

	// Initialize the Server Dispatcher
	mux := http.NewServeMux()

	// Create a file server to serve static files
	fs := http.FileServer(http.Dir("/app/assets"))

	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Ask the dispatcher to handle the root path
	mux.HandleFunc("/", indexHandler)

	// Start the server using http package and set the port and dispatcher
	fmt.Println("Listening on port " + os.Getenv("PORT"))
	http.ListenAndServe(":"+port, mux)
}
