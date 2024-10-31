package main

import (
	"library/config"
	"log"
	"net/http"
)

func main() {
	// Initialize the database
	db, err := config.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the router
	r := router.NewRouter(db)

	// Start the server
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
