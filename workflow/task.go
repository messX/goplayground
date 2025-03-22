package workflow

import (
	"log"
)

type Task struct {
	Name              string
	Execute           func() error
	RetryCount        int
	Status            string
	MAX_FAILURE_LIMIT int
	Dependency        []*Task
	IS_Done           chan bool
}

func (task *Task) ExecuteTask() error {
	// wait for all dependencies to be done
	for _, dep := range task.Dependency {
		<-dep.IS_Done
	}
	log.Println("Executing task: ", task.Name)
	// execute the task
	if task.RetryCount > task.MAX_FAILURE_LIMIT {
		task.Status = "Failed"
		return nil
	}
	task.Status = "Running"
	err := task.Execute()
	if err != nil {
		task.RetryCount++
		return task.ExecuteTask()
	}
	task.Status = "Completed"
	log.Println("Task completed: ", task.Name)
	close(task.IS_Done)
	return nil
}
