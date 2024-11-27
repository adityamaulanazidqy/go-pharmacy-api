package main

import (
	"go-pharmacy-api/config"
	"go-pharmacy-api/routes"
	"log"
	"net/http"
)

func main() {
	// Hubungkan ke database
	config.ConnectDatabase()
	defer config.DB.Close()

	// Konfigurasi routing
	router := routes.RegisterRoutes()

	// Jalankan server
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
