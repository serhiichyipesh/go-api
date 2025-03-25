package main

import (
	"github.com/joho/godotenv"
	"github.com/serhiichyipesh/go-api/internal/env"
	"github.com/serhiichyipesh/go-api/internal/store"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config{
		addr: env.GetString("PORT", ":8080"),
	}

	storage := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  storage,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
