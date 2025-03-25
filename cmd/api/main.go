package main

import (
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/serhiichyipesh/go-api/internal/env"
	"github.com/serhiichyipesh/go-api/internal/store"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("postgres", env.GetString("DB_ADDR", ""))
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	cfg := config{
		addr: env.GetString("PORT", ":8080"),
	}

	storage := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  storage,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
