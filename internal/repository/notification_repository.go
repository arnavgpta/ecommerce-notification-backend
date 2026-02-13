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

func (r *NotificationRepository) ExistsPendingNotification(
	ctx context.Context,
	userID int,
	notificationType string,
) (bool, error) {

	query := `
		SELECT COUNT(1)
		FROM notifications
		WHERE user_id = $1
		AND type = $2
		AND status = 'pending'
	`

	var count int
	err := r.db.QueryRowContext(ctx, query,
		userID,
		notificationType,
	).Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil
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
