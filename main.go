package main

// go mod tidy updates the go.mod file with all currently used dependencies
// go mod init is used when making an new module to make the go.mod file
// go mod vendor makes the vendor dir in workspace
// good for a self-contaned project and includes all dependencies

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not defined in env")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	Router1 := chi.NewRouter()

	Router1.HandleFunc("/ready", handlerReadiness)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	// print format like C
	log.Printf("Server starting on port %v", portString)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PORT:", portString)
}
