package pubsub

import (
	"fmt"

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
		fmt.Printf("I am topic tp1 and msg : %s", msg)
	}
	f2 := func(msg string) {
		fmt.Printf("I am topic tp2 and msg : %s", msg)
	}

	ag.Subscribe(tp1, f1)
	ag.Subscribe(tp2, f2)

	ag.Publish(tp1, "test 1")
	ag.Publish(tp2, "test 2")

	ag.Close()
}
