package core_test

import (
	"testing"

	"github.com/iamcc/ddd-go/core"
)

type testAggregateRoot struct {
	core.AggregateRoot
	id string
}

type testDomainEvent struct {
	id string
}

func (testDomainEvent) Name() string {
	return "testDomainEvent"
}

func TestDomainEvent(t *testing.T) {
	ar := &testAggregateRoot{id: "id"}
	ar.AddEvent(&testDomainEvent{ar.id})
	handled := false
	core.DefaultDomainEventPublisher.Subscribe(testDomainEvent{}, core.DomainEventHandler(func(e any) {
		handled = true
		ee, ok := e.(*testDomainEvent)
		if !ok {
			t.Error("event should be testDomainEvent")
		}
		if ee.id != ar.id {
			t.Error("event id should be equal to aggregate root id")
		}
	}))
	core.DefaultDomainEventPublisher.PublishAll(ar.Events())
	if !handled {
		t.Error("event should be handled")
	}
}
