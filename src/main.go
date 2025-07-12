package main

import (
	"net/http"
	"rolerocket/logger"
	"rolerocket/routes"
)

func main() {
	router := routes.Routes()
	logger.Main()

	server := http.Server{
		//Addr: ":8080",
		Addr:    "localhost:8080", // ? use localhost:xxxx to make it not ask for admin permissions
		Handler: logger.Middlware(router),
	}
	server.ListenAndServe()
}
