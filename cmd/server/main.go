package main

import (
	"net/http"
	sqlite "rolerocket/internal/db"
	"rolerocket/internal/logger"
	"rolerocket/internal/routes"
)

func main() {
	router := routes.Routes()
	db := sqlite.Init()
	logger.Main(db)

	// ? use localhost:xxxx to make it not ask for admin permissions
	// ? use :8080 for production
	port := "localhost:8080"
	logger.Info("Starting server on port: " + port)
	server := http.Server{
		Addr:    port,
		Handler: logger.Middlware(router),
	}
	server.ListenAndServe()

}
