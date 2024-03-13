package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

func runWorker() {
	temporalClient, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Message temporal client", err)
		return
	}
	defer temporalClient.Close()

	myWorker := worker.New(temporalClient, queueName, worker.Options{})
	myWorker.RegisterWorkflow(SimpleWorkflowDefinition)
	myWorker.RegisterActivity(SimpleActivity)
	err = myWorker.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to run Message temporal worker", err)
	}
}

type WorkflowParam struct {
	Message string
	Code    int
}

type WorkflowResult struct {
	Value string
}

func SimpleWorkflowDefinition(ctx workflow.Context, param WorkflowParam) (WorkflowResult, error) {
	logger := workflow.GetLogger(ctx)
	if param.Code <= 0 {
		return WorkflowResult{}, errors.New(fmt.Sprint("Forbidden Code", param.Code))
	}

	logger.Info("Param ", param.Message, param.Code)

	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	})
	var result string
	err := workflow.ExecuteActivity(ctx, SimpleActivity, param.Message).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity execution failed:", err)
		return WorkflowResult{}, err
	}
	logger.Info("Activity execution result:", result)
	return WorkflowResult{Value: "workflow completed"}, nil
}

func SimpleActivity(ctx context.Context, message string) (string, error) {
	log.Println("SimpleActivity: " + message)
	return "activity completed", nil
}
