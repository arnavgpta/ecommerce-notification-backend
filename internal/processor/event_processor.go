package processor

import (
	"context"
	"log"
	"time"

	"github.com/arnavgpta/ecommerce-notification-backend/internal/models"
	"github.com/arnavgpta/ecommerce-notification-backend/internal/repository"
	"github.com/arnavgpta/ecommerce-notification-backend/internal/rules"
)

type EventProcessor struct {
	queue            chan models.CreateEventRequest
	notificationRepo *repository.NotificationRepository
}

func NewEventProcessor(
	bufferSize int,
	notificationRepo *repository.NotificationRepository,
) *EventProcessor {
	return &EventProcessor{
		queue:            make(chan models.CreateEventRequest, bufferSize),
		notificationRepo: notificationRepo,
	}
}

func (p *EventProcessor) StartWorker() {
	go func() {
		for event := range p.queue {
			p.handleEvent(event)
		}
	}()
}

func (p *EventProcessor) Enqueue(event models.CreateEventRequest) {
	p.queue <- event
}

func (p *EventProcessor) handleEvent(event models.CreateEventRequest) {

	rule := rules.DetermineNotification(event.EventType)

	if !rule.ShouldNotify {
		return
	}

	if rule.Delay > 0 {
		go p.scheduleNotification(event.UserID, rule.NotificationType, rule.Delay)
		return
	}

	p.createNotification(event.UserID, rule.NotificationType)
}

func (p *EventProcessor) scheduleNotification(
	userID int,
	notificationType string,
	delay time.Duration,
) {
	time.Sleep(delay)

	p.createNotification(userID, notificationType)
}

func (p *EventProcessor) createNotification(
	userID int,
	notificationType string,
) {
	err := p.notificationRepo.CreateNotification(
		context.Background(),
		userID,
		notificationType,
	)

	if err != nil {
		log.Printf("failed to create notification: %v", err)
		return
	}

	log.Printf("notification created: %s for user %d",
		notificationType,
		userID,
	)
}
