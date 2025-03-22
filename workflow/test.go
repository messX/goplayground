package workflow

import (
	"log"
	"time"
)

func TestWorkflow() {
	log.Println("Starting workflow")
	t1 := &Task{
		Name: "Task 1",
		Execute: func() error {
			log.Println("Executing Task 1")
			time.Sleep(1 * time.Second)
			return nil
		},
		RetryCount:        3,
		Status:            "Not Started",
		MAX_FAILURE_LIMIT: 3,
		Dependency:        []*Task{},
		IS_Done:           make(chan bool),
	}

	t2 := &Task{
		Name: "Task 2",
		Execute: func() error {
			log.Println("Executing Task 2")
			time.Sleep(1 * time.Second)
			return nil
		},
		RetryCount:        3,
		Status:            "Not Started",
		MAX_FAILURE_LIMIT: 3,
		Dependency:        []*Task{t1},
		IS_Done:           make(chan bool),
	}

	t3 := &Task{
		Name: "Task 3",
		Execute: func() error {
			log.Println("Executing Task 3")
			time.Sleep(1 * time.Second)
			return nil
		},
		RetryCount:        3,
		Status:            "Not Started",
		MAX_FAILURE_LIMIT: 3,
		Dependency:        []*Task{t2},
		IS_Done:           make(chan bool),
	}

	w := &Workflow{
		Tasks:  []*Task{t1, t2, t3},
		Name:   "Workflow 1",
		Status: "Not Started",
	}

	err := w.ExecuteWorkflow()
	if err != nil {
		log.Println("Error in executing workflow")
	}
	log.Println("Workflow completed", w.Status)
}
