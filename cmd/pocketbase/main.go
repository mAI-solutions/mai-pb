package main

import (
	"log"
	"pb-mockup/internal/handlers"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
    "github.com/profclems/go-dotenv"
)

func main() {
    app := pocketbase.New()

    err := dotenv.LoadConfig()
    if err != nil {
        log.Fatalf("error loading .env file: %v", err)
    }

    // serves static files from the provided public dir (if exists)
    app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
        e.Router.POST("/feed", handlers.FeedHandler)
        e.Router.POST("/complet", handlers.CompĺetionHandler)
        return nil
    })

    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
}

