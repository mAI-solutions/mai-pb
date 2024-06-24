package main

import (
	"log"
	"pb-mockup/internal/handlers"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
    app := pocketbase.New()

    // serves static files from the provided public dir (if exists)
    app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
        e.Router.POST("/feed", handlers.TestHandler)
        return nil
    })

    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
}

