package repository

import (
	"context"
	"database/sql"
)

type NotificationRepository struct {
	db *sql.DB
}

func NewNotificationRepository(db *sql.DB) *NotificationRepository {
	return &NotificationRepository{db: db}
}

func (r *NotificationRepository) CreateNotification(
	ctx context.Context,
	userID int,
	notificationType string,
) error {

	query := `
		INSERT INTO notifications (user_id, type)
		VALUES ($1, $2)
	`

	_, err := r.db.ExecContext(ctx, query, userID, notificationType)
	return err
}
