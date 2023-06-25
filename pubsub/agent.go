package pubsub

import (
	"sync"

	log "github.com/sirupsen/logrus"
)

type Ipubsub interface {
	Subscribe(string) <-chan string
	Pulish(string, string)
}

type Agent struct {
	mu     sync.Mutex
	subs   map[string][]chan string
	quit   chan struct{}
	closed bool
}

func NewAgent() *Agent {
	return &Agent{
		subs: make(map[string][]chan string),
		quit: make(chan struct{}),
	}
}

func (a *Agent) Subscribe(topic string, handler func(m string)) chan string {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.closed {
		log.Debug("Agent has stopped pub-sub")
	}
	ch := make(chan string)
	a.subs[topic] = append(a.subs[topic], ch)
	go func() {
		for _, c := range a.subs[topic] {
			msg := <-c
			handler(msg)
		}

	}()
	log.Infof("Subscriber added to topic %s", topic)
	return ch
}

func (a *Agent) Publish(topic string, msg string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.closed {
		log.Debug("Agent has stopped pub-sub")
	}
	for _, ch := range a.subs[topic] {
		ch <- msg
	}
	log.Infof("Successfully published in topic: %s", topic)
}

func (a *Agent) Close() {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.closed {
		log.Debug("Agent has stopped pub-sub")
	}
	if a.closed {
		return
	}
	a.closed = true
	close(a.quit)
	for _, chs := range a.subs {
		for _, ch := range chs {
			close(ch)
		}
	}
	log.Debug("All channels closed quitting!")
}
