package processor

import (
	"log"

	"github.com/arnavgpta/ecommerce-notification-backend/internal/models"
)

type EventProcessor struct {
	queue chan models.CreateEventRequest
}

func NewEventProcessor(bufferSize int) *EventProcessor {
	return &EventProcessor{
		queue: make(chan models.CreateEventRequest, bufferSize),
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
	log.Printf("Processing event: %s for user %d",
		event.EventType,
		event.UserID,
	)

}
