package main

import (
 	"log"
 	"net/http"
 	"os"
	"github.com/gorilla/handlers"
	"/router"
)

func main() {
	// Get the "PORT" env variable
    // port := os.Getenv("PORT")
    
    port := "80"

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := router.NewRouter() // create routes

	// These two lines are important if you're designing a front-end to utilise this API methods
	allowedOrigins := handlers.AllowedOrigins([]string{"*"}) 
 	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	// Launch server with CORS validations
 	log.Fatal(http.ListenAndServe(":" + port, handlers.CORS(allowedOrigins, allowedMethods)(router)))
}