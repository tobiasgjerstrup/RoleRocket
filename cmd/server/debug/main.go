package main

import (
	"context"
	"net/http"
	sqlite "rolerocket/internal/db"
	"rolerocket/internal/logger"
	"rolerocket/internal/routes"
)

func main() {
	router := routes.Routes(true)
	db := sqlite.Init()
	logger.Main(db)

	// ? use localhost:xxxx to make it not ask for admin permissions
	// ? use :8080 for production and benchmarking! This changes how TCP behaves so it's very important to use :8080
	port := "localhost:8080"
	logger.Info(context.Background(), "Starting server on port: "+port)
	server := http.Server{
		Addr:    port,
		Handler: logger.Middlware(router),
	}
	server.ListenAndServe()

}
