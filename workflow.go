package main

import (
	"fmt"
	"go.temporal.io/sdk/workflow"
	"sort"
)

const queueName = "work-queue"

type WorkflowParam struct {
	Message string
	Size    int
}

type WorkflowResult struct {
	Value string
}

func ParentWorkflow(ctx workflow.Context, param WorkflowParam) (WorkflowResult, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("Workflow started, param:", param.Message)

	childFutures := make([]workflow.Future, 0)
	for i := 0; i < param.Size; i++ {
		future := workflow.ExecuteChildWorkflow(ctx, ChildWorkflow, i)
		childFutures = append(childFutures, future)
	}

	childResults := make([]int, 0)
	for _, future := range childFutures {
		var result int
		if err := future.Get(ctx, &result); err != nil {
			return WorkflowResult{}, err
		}
		childResults = append(childResults, result)
	}
	sort.Ints(childResults)

	result := fmt.Sprint("Completed ", childResults[0], childResults[len(childResults)-1])
	logger.Info("Parent workflow execution: " + param.Message + " " + result)
	return WorkflowResult{Value: result}, nil
}

func ChildWorkflow(ctx workflow.Context, number int) (int, error) {
	logger := workflow.GetLogger(ctx)
	greeting := fmt.Sprint("Hello Child", number, "!")
	logger.Info("Child workflow execution:" + greeting)
	return number, nil
}
