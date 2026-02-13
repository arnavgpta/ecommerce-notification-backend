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

func (r *EventRepository) HasRecentOrder(
	ctx context.Context,
	userID int,
) (bool, error) {

	query := `
		SELECT COUNT(1)
		FROM events
		WHERE user_id = $1
		AND event_type = 'order_placed'
		AND created_at > NOW() - INTERVAL '5 minutes'
	`

	var count int
	err := r.db.QueryRowContext(ctx, query, userID).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
