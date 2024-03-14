package main

import (
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
)

func runWorker(temporalClient client.Client) {
	myWorker := worker.New(temporalClient, queueName, worker.Options{})
	myWorker.RegisterWorkflow(ParentWorkflow)
	myWorker.RegisterWorkflow(ChildWorkflow)
	err := myWorker.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to run a temporal worker", err)
	}
	log.Println("Temporal worker started")
}
