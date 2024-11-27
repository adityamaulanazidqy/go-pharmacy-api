package main

import (
	"go-pharmacy-api/config"
	"go-pharmacy-api/routes"
	"log"
	"net/http"
)

func main() {
	config.ConnectDatabase()
	defer config.DB.Close()

	router := routes.RegisterRoutes()

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
