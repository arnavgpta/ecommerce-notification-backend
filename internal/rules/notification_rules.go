package rules

import "github.com/arnavgpta/ecommerce-notification-backend/internal/models"

func DetermineNotification(event models.CreateEventRequest) (string, bool) {

	switch event.EventType {

	case "added_to_cart":
		return "cart_reminder", true

	case "order_placed":
		return "order_confirmation", true

	case "user_signed_up":
		return "welcome_message", true

	default:
		return "", false
	}
}
