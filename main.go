package main

// go mod tidy updates the go.mod file with all currently used dependencies
// go mod init is used when making an new module to make the go.mod file
// go mod vendor makes the vendor dir in workspace
// good for a self-contaned project and includes all dependencies

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dukebward/news_grabber/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	// underscore is to include even if not using
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not defined in env")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not defined in env")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database", err)
	}

	db := database.New(conn)
	apiCfg := apiConfig{
		DB: db,
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

	//something like localhost:8000/v1/healthz
	//the Get method only allows Get requests
	// anything that uses the mid auth function is an authenticated endpoint
	Router1.Get("/healthz", handlerReadiness)
	Router1.Get("/error", handlerError)
	Router1.Post("/users", apiCfg.handlerCreateUser)
	Router1.Get("/users", apiCfg.midAuth(apiCfg.handlerGetUser))
	Router1.Post("/feeds", apiCfg.midAuth(apiCfg.handlerCreateFeed))
	Router1.Get("/feeds", apiCfg.handlerGetFeeds)
	Router1.Post("/feed_follows", apiCfg.midAuth(apiCfg.handlerCreateFeedFollow))
	Router1.Get("/feed_follows", apiCfg.midAuth(apiCfg.handlerGetFeedFollow))
	Router1.Delete("/feed_follows/{feedFollowID}", apiCfg.midAuth(apiCfg.handlerDeleteFeedFollow))
	Router1.Get("/posts", apiCfg.midAuth(apiCfg.handlerGetPostsForUser))

	router.Mount("/v1", Router1)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	go startScraping(db, 10, time.Minute)

	// print format like C
	log.Printf("Server serving on port %v", portString)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PORT:", portString)
}
