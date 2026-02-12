package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/arnavgpta/ecommerce-notification-backend/internal/models"
	"github.com/arnavgpta/ecommerce-notification-backend/internal/processor"
	"github.com/arnavgpta/ecommerce-notification-backend/internal/repository"
)

type EventHandler struct {
	repo      *repository.EventRepository
	processor *processor.EventProcessor
}

func NewEventHandler(
	repo *repository.EventRepository,
	processor *processor.EventProcessor,
) *EventHandler {
	return &EventHandler{
		repo:      repo,
		processor: processor,
	}
}

func (h *EventHandler) IngestEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.CreateEventRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if req.UserID == 0 || req.EventType == "" {
		http.Error(w, "user_id and event_type required", http.StatusBadRequest)
		return
	}

	err := h.repo.CreateEvent(r.Context(), req)
	if err != nil {
		http.Error(w, "could not store event", http.StatusInternalServerError)
		return
	}

	h.processor.Enqueue(req)

	w.WriteHeader(http.StatusCreated)
}
