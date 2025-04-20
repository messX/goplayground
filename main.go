package main

import (
	//"github.com/messx/goplayground/pubsub"
	//"github.com/messx/goplayground/lru"
	//expobackoff "github.com/messx/goplayground/expo_backoff"
	//"github.com/messx/goplayground/workflow"
	"github.com/messx/goplayground/webhook"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"user": "admin",
	}).Info("Some interesting info")
	//pubsub.Test()
	//lru.Test()
	//expobackoff.Test()
	//workflow.TestWorkflow()
	webhook.Run()
}
