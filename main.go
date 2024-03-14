package main

import (
	"go.temporal.io/sdk/client"
	"log"
	"os"
)

func main() {
	temporalClient, err := client.Dial(client.Options{
		HostPort: os.Getenv("HOSTPORT"),
	})
	if err != nil {
		log.Fatalln("Unable to create a temporal client", err)
	}
	defer temporalClient.Close()

	role := os.Getenv("MW_ROLE")
	if role == "leader" {
		log.Println("I am leader.")
		runLeader(temporalClient)
		runWorker(temporalClient)
		return
	} else if role == "worker" {
		log.Println("I am worker.")
		runWorker(temporalClient)
		return
	}
	panic("Unknown run mode")
}
