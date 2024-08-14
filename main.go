package main

import (
	"gofr.dev/pkg/gofr"

	"find_pairs/handlers"
)

func main() {
	// Initial a new gofr application
	app := gofr.New()

	// Add all the routes
	app.POST("/find-pairs", handlers.FindPairs)

	// Start the application
	app.Run()
}
