package main

import (
	configservers "admin-go/configServers"
	routers "admin-go/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func main() {
	configservers.StartUp()
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Accept", "User-Agent", "Origin", "auth_token", "X-CSRF-Token"})
	methods := handlers.AllowedMethods([]string{"POST", "GET", "OPTIONS", "PUT", "DELETE", "PATCH"})
	origins := handlers.AllowedOrigins([]string{"*"})

	router := routers.InitRoutes()
	server := &http.Server{
		Addr:    os.Getenv("SERVER"),
		Handler: router,
	}

	log.Println("Listening 5000...")
	http.ListenAndServe(server.Addr, handlers.CORS(headers, methods, origins)(router))
}
