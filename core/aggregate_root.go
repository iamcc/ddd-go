package core

type IAggregateRoot interface {
	AddEvent(DomainEvent)
	Events() []DomainEvent
	ClearEvents()
}

type AggregateRoot struct {
	events []DomainEvent
}

var _ IAggregateRoot = &AggregateRoot{}

// AddEvent implements IAggregateRoot.
func (ar *AggregateRoot) AddEvent(e DomainEvent) {
	ar.events = append(ar.events, e)
}

// ClearEvents implements IAggregateRoot.
func (ar *AggregateRoot) ClearEvents() {
	ar.events = []DomainEvent{}
}

// Events implements IAggregateRoot.
func (ar *AggregateRoot) Events() []DomainEvent {
	return ar.events
}
