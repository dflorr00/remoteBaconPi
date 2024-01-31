package main

import (
	handlers "backend/Handlers"
	tools "backend/Tools"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
)

func main() {

	tools.Discovery()
	go tools.ReloadStatus30s()

	mux := http.NewServeMux()
	mux.HandleFunc("/signIn", handlers.HandlerSignIn)
	mux.HandleFunc("/devices/", handlers.HandlerDevices)
	mux.HandleFunc("/users/", handlers.HandlerUsers)
	mux.HandleFunc("/views/", handlers.HandlerViews)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8001", "http://localhost:8080"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "Custom-Header", "Cookie"},
		AllowCredentials: true,
	})
	handler := c.Handler(mux)
	port := ":8080"
	log.Printf("Iniciando servidor en el puerto %s...\n", port)

	http.ListenAndServe(port, handler)
}
