package main

import (
	"github.com/messx/goplayground/pubsub"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"user": "admin",
	}).Info("Some interesting info")
	pubsub.Test()
}
