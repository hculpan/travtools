package main

import (
	"embed"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/hculpan/travtools/handlers"
	"github.com/joho/godotenv"
)

// content holds our static web server content.
//
//go:embed assets/* templates/*
var content embed.FS

func routes() {
	// Serve static files from "assets" directory
	http.Handle("/assets/", http.FileServer(http.FS(content)))

	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/message", handlers.MessagePage)
	http.HandleFunc("/planet-generator", handlers.PlanetGeneratorPage)

	http.HandleFunc("/trade-generator", handlers.TradeGeneratorPage)
}

func main() {
	// Load the environment variables from .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	handlers.SetContent(&content)
	routes()

	log.Printf("Server started on port %s\n", port)
	http.ListenAndServe(":"+port, LogRequest(http.DefaultServeMux))
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s %v", r.Method, r.RequestURI, r.RemoteAddr, time.Since(startTime))
	})
}
