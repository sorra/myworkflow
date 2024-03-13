package main

import (
	"os"
)

const queueName = "work-queue"

func main() {
	for _, arg := range os.Args {
		if arg == "leader" {
			runLeader()
			return
		} else if arg == "worker" {
			runWorker()
			return
		}
	}
	panic("Unknown run mode")
}
