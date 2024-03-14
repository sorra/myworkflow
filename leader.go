package main

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"
)

func runLeader() {
	temporalClient, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create a temporal client", err)
	}
	defer temporalClient.Close()

	workflowOptions := client.StartWorkflowOptions{
		TaskQueue: queueName,
	}
	param := WorkflowParam{
		Message: "hello world",
		Code:    1,
	}
	workflowRun, err := temporalClient.ExecuteWorkflow(
		context.Background(), workflowOptions,
		SimpleWorkflowDefinition, param)
	if err != nil {
		log.Fatalln("Unable to get workflow run:", err)
	}
	log.Println("Workflow RunID:", workflowRun.GetRunID())

	var result WorkflowResult
	err = workflowRun.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Workflow execution failure:", err)
	}
	log.Println("Work execution result:", result.Value)
}
