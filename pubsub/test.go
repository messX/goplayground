package pubsub

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

func Test() {
	log.WithFields(log.Fields{
		"user": "admin",
	}).Info("Hello lets begin!")
	ag := NewAgent()
	tp1 := "tp1"
	tp2 := "tp2"
	f1 := func(msg string) {
		time.Sleep(100)
		fmt.Printf("I am topic tp1 and msg : %s", msg)
	}
	f2 := func(msg string) {
		fmt.Printf("I am topic tp2 and msg : %s", msg)
	}

	ag.Subscribe(tp1, f1)
	ag.Subscribe(tp2, f2)

	ag.Publish(tp1, "test 1")
	ag.Publish(tp2, "test 2")

	// Wait for goroutines to finish
	// wg := sync.WaitGroup{}
	// wg.Add(2)

	/*
		Todo:
		1. Add indefinite listening
		2. Add subscriber module to handle the execution as per flow during runtime.
	*/
	ag.Close()
}
