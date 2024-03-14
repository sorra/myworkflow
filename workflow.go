package main

import (
	"errors"
	"fmt"
	"go.temporal.io/sdk/workflow"
	"sort"
)

type WorkflowParam struct {
	Message string
	Code    int
}

type WorkflowResult struct {
	Value string
}

func ParentWorkflow(ctx workflow.Context, param WorkflowParam) (WorkflowResult, error) {
	logger := workflow.GetLogger(ctx)
	if param.Code <= 0 {
		return WorkflowResult{}, errors.New(fmt.Sprint("Forbidden Code", param.Code))
	}

	logger.Info("Workflow started, param:", param.Message, param.Code)

	childFutures := make([]workflow.Future, 0)
	for i := 0; i < 200; i++ {
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
	return WorkflowResult{Value: result}, nil
}

func ChildWorkflow(ctx workflow.Context, number int) (int, error) {
	logger := workflow.GetLogger(ctx)
	greeting := fmt.Sprint("Hello Child", number, "!")
	logger.Info("Child workflow execution: " + greeting)
	return number, nil
}
