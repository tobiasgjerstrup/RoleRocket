package main

import (
	"net/http"
	"rolerocket/databases/sqlite"
	"rolerocket/logger"
	"rolerocket/routes"
)

func main() {
	router := routes.Routes()
	logger.Main()
	sqlite.Main()

	// ? use localhost:xxxx to make it not ask for admin permissions
	// ? use :8080 for production
	port := "localhost:8080"
	logger.Logger.Info("Starting server on port: " + port)
	server := http.Server{
		Addr:    port,
		Handler: logger.Middlware(router),
	}
	server.ListenAndServe()

}
