package main

import (
	"log"
	"time"

	"github.com/dukebward/news_grabber/internal/database"
)

// scrapers run in background constantly
func startScraping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Printf("Scraping on %v goroutines every %s duration", concurrency, timeBetweenRequest)

}
