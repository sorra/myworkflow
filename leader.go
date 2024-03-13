package main

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"
)

func runLeader() {
	temporalClient, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Message temporal client", err)
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
		return
	}
	log.Println("Workflow RunID:", workflowRun.GetRunID())
	var result WorkflowResult
	err = workflowRun.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Workflow execution failure:", err)
		return
	}
	log.Println("Work execution result:", result.Value)
}
