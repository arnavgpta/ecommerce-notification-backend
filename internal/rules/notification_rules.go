package rules

import "time"

type RuleResult struct {
	NotificationType string
	Delay            time.Duration
	ShouldNotify     bool
}

func DetermineNotification(eventType string) RuleResult {

	switch eventType {

	case "added_to_cart":
		return RuleResult{
			NotificationType: "cart_reminder",
			Delay:            30 * time.Second,
			ShouldNotify:     true,
		}

	case "order_placed":
		return RuleResult{
			NotificationType: "order_confirmation",
			Delay:            0,
			ShouldNotify:     true,
		}

	case "user_signed_up":
		return RuleResult{
			NotificationType: "welcome_message",
			Delay:            0,
			ShouldNotify:     true,
		}

	default:
		return RuleResult{
			ShouldNotify: false,
		}
	}
}
