package main

import (
	//"github.com/messx/goplayground/pubsub"
	"github.com/messx/goplayground/lru"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"user": "admin",
	}).Info("Some interesting info")
	//pubsub.Test()
	lru.Test()
}
