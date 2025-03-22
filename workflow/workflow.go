package workflow

import (
	"log"
	"sync"
)

type Workflow struct {
	Tasks  []*Task
	Name   string
	Status string
	WG     sync.WaitGroup // Added to fix the error: undefined: sync in workflow/workflow.go
}

func (workflow *Workflow) ExecuteWorkflow() error {
	workflow.Status = "Running"
	for _, task := range workflow.Tasks {
		workflow.WG.Add(1) // Added to fix the error: undefined: sync in workflow/workflow.go
		go func(task *Task) {
			defer workflow.WG.Done() // Added to fix the error: undefined: sync in workflow/workflow.go
			err := task.ExecuteTask()
			if err != nil {
				return
			}
		}(task)
	}
	workflow.WG.Wait() // Added to fix the error: undefined: sync in workflow/workflow.go
	log.Println("Workflow completed")
	workflow.Status = "Completed"
	return nil
}
