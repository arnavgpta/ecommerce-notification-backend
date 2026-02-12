package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/arnavgpta/ecommerce-notification-backend/internal/handlers"
	"github.com/arnavgpta/ecommerce-notification-backend/internal/repository"
	_ "github.com/lib/pq"

	"github.com/arnavgpta/ecommerce-notification-backend/internal/processor"
)

func main() {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("cannot connect to DB:", err)
	}

	mux := http.NewServeMux()

	eventRepo := repository.NewEventRepository(db)
	notificationRepo := repository.NewNotificationRepository(db)

	eventProcessor := processor.NewEventProcessor(
		100,
		notificationRepo,
	)

	eventProcessor.StartWorker()

	eventHandler := handlers.NewEventHandler(
		eventRepo,
		eventProcessor,
	)

	mux.HandleFunc("/events", eventHandler.IngestEvent)

	log.Println("server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
