package notifications

import (
	"container/list"
	"sync"

	events "github.com/docker/go-events"
)

type eventQueue struct {
	sink      events.Sink
	events    *list.List
	listeners []eventQueueListener
	cond      *sync.Cond
	mu        sync.Mutex
	closed    bool
}

type eventQueueListener interface {
	ingress(event events.Event)
	egress(event events.Event)
}

func newEventQueue(sink events.Sink, listeners ...eventQueueListener) *eventQueue {
	_ = "STUB: not implemented"
	return nil
}

func (eq *eventQueue) Write(event events.Event) error { _ = "STUB: not implemented"; return nil }

func (eq *eventQueue) Close() error { _ = "STUB: not implemented"; return nil }

func (eq *eventQueue) run() { _ = "STUB: not implemented"; return }

func (eq *eventQueue) next() events.Event { _ = "STUB: not implemented"; return *new(events.Event) }

type ignoredSink struct {
	events.Sink
	ignoreMediaTypes map[string]bool
	ignoreActions    map[string]bool
}

func newIgnoredSink(sink events.Sink, ignored []string, ignoreActions []string) events.Sink {
	_ = "STUB: not implemented"
	return *new(events.Sink)
}

func (imts *ignoredSink) Write(event events.Event) error { _ = "STUB: not implemented"; return nil }

func (imts *ignoredSink) Close() error { _ = "STUB: not implemented"; return nil }
