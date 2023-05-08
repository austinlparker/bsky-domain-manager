package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Register HTTP handlers for different routes
	http.HandleFunc("/", handleRoot)          // Handle requests to the root path
	http.HandleFunc("/get/", handleGet)       // Handle requests to get data
	http.HandleFunc("/create/", handleCreate) // Handle requests to create data
	http.HandleFunc("/update/", handleUpdate) // Handle requests to update data
	http.HandleFunc("/xrpc/com.atproto.identity.resolveHandle", handleResolveHandle)

	// Serve static files (JS, CSS, etc.)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start the HTTP server with the specified port
	log.Printf("Starting server on port %s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
