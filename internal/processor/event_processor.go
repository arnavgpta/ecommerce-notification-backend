package processor

import (
	"context"
	"log"

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

	notificationType, shouldNotify :=
		rules.DetermineNotification(event)

	if !shouldNotify {
		return
	}

	err := p.notificationRepo.CreateNotification(
		context.Background(),
		event.UserID,
		notificationType,
	)

	if err != nil {
		log.Printf("failed to create notification: %v", err)
		return
	}

	log.Printf("notification created: %s for user %d",
		notificationType,
		event.UserID,
	)
}
