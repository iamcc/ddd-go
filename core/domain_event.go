package core

type DomainEvent interface {
	Name() string
}

type DomainEventHandler func(any)

type DomainEventPublisher interface {
	Publish(DomainEvent)
	PublishAll([]DomainEvent)
	Subscribe(DomainEvent, DomainEventHandler)
}

var DefaultDomainEventPublisher = &defaultDomainEventPublisher{
	handlers: map[string][]DomainEventHandler{},
}

type defaultDomainEventPublisher struct {
	handlers map[string][]DomainEventHandler
}

var _ DomainEventPublisher = &defaultDomainEventPublisher{}

// Publish implements DomainEventPublisher.
func (p *defaultDomainEventPublisher) Publish(e DomainEvent) {
	handlers := p.handlers[e.Name()]
	for _, h := range handlers {
		h(e)
	}
}

// PublishAll implements DomainEventPublisher.
func (p *defaultDomainEventPublisher) PublishAll(events []DomainEvent) {
	for _, e := range events {
		p.Publish(e)
	}
}

// Subscribe implements DomainEventPublisher.
func (p *defaultDomainEventPublisher) Subscribe(e DomainEvent, h DomainEventHandler) {
	handlers := p.handlers[e.Name()]
	handlers = append(handlers, h)
	p.handlers[e.Name()] = handlers
}
