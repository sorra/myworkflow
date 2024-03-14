package main

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"
)

func runLeader(temporalClient client.Client) {
	workflowId := "cron-1m"
	lastWorkflow := temporalClient.GetWorkflow(context.Background(), workflowId, "")
	if lastWorkflow != nil {
		err := temporalClient.CancelWorkflow(context.Background(), workflowId, "")
		if err != nil {
			log.Println("Unable to cancel the last workflow:", workflowId, err)
		} else {
			log.Println("Cancelled last workflow:", workflowId)
		}
	}

	workflowOptions := client.StartWorkflowOptions{
		ID:           workflowId,
		TaskQueue:    queueName,
		CronSchedule: "* * * * *",
	}
	param := WorkflowParam{
		Message: "Hello " + workflowId,
		Size:    10,
	}
	workflowRun, err := temporalClient.ExecuteWorkflow(
		context.Background(), workflowOptions,
		ParentWorkflow, param)
	if err != nil {
		log.Fatalln("Unable to get workflow run:", err)
	}
	log.Println("Workflow", workflowId, "started, RunID:", workflowRun.GetRunID())
}
