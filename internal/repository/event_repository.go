package repository

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/arnavgpta/ecommerce-notification-backend/internal/models"
)

type EventRepository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (r *EventRepository) CreateEvent(
	ctx context.Context,
	req models.CreateEventRequest,
) error {

	metaJSON, err := json.Marshal(req.Metadata)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO events (user_id, event_type, metadata)
		VALUES ($1, $2, $3)
	`

	_, err = r.db.ExecContext(ctx, query,
		req.UserID,
		req.EventType,
		metaJSON,
	)

	return err
}
