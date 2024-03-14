package main

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"
)

func runLeader(temporalClient client.Client) {
	executeCronWorkflow(temporalClient, "cron-1m", "* * * * *")
	executeCronWorkflow(temporalClient, "cron-5m", "*/5 * * * *")
}

func executeCronWorkflow(temporalClient client.Client, workflowId string, cronSchedule string) {
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
		CronSchedule: cronSchedule,
	}
	param := WorkflowParam{
		Message: workflowId,
		Size:    200,
	}
	workflowRun, err := temporalClient.ExecuteWorkflow(
		context.Background(), workflowOptions,
		ParentWorkflow, param)
	if err != nil {
		log.Fatalln("Unable to get workflow run:", err)
	}
	log.Println("Workflow", workflowId, "started, RunID:", workflowRun.GetRunID())
}
