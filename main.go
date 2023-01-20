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

indexPath := "/app/index.html"

if _, err := os.Stat(indexPath); os.IsNotExist(err) {
	indexPath = "index.html"
}

var tpl = template.Must(template.ParseFiles(indexPath))

// Create the handler for the root index page
func indexHandler(w http.ResponseWriter, r *http.Request) {

	// This template is executed using html/template package
	tpl.Execute(w, nil)
}

func main() {

	// Load the .env file in current directory
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file.")
	}

	// Get the port from the environment variable
	port := os.Getenv("PORT")

	// Initialize the Server Dispatcher
	mux := http.NewServeMux()

	// Create a file server to serve static files
	fs := http.FileServer(http.Dir("assets"))

	// Handle the static files using the file server
	assetsPath := "/app/assets/"
	if _, err := os.Stat(assetsPath); os.IsNotExist(err) {
		assetsPath = "assets/"
	}

	mux.Handle("/assets/", http.StripPrefix(assetsPath, fs))

	// Ask the dispatcher to handle the root path
	mux.HandleFunc("/", indexHandler)

	// Start the server using http package and set the port and dispatcher
	fmt.Println("Listening on port " + os.Getenv("PORT"))
	http.ListenAndServe(":"+port, mux)
}
