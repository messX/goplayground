package expobackoff

import (
	"fmt"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
)

const MAX_FAILURE_LIMIT = 10

type Expobackoff struct {
	CurrentBackoffCount int
	TaskChan            chan string
}

func (expobackoff *Expobackoff) dummy_backoff(is_error bool) (bool, error) {
	log.Debug("Dummy backoff method")
	if is_error {
		return false, fmt.Errorf("Sending error")
	}
	return true, nil
}

/*
simple_backoff implements a simple backoff mechanism. It logs the process and
returns an error if the maximum failure limit is exceeded.
*/
func (expobackoff *Expobackoff) simple_backoff() error {
	log.Debug("Initiating simple_backoff function call")

	// Check if the current backoff count exceeds the maximum failure limit
	if expobackoff.CurrentBackoffCount > MAX_FAILURE_LIMIT {
		log.Errorf("Error limit exceeded: %d", MAX_FAILURE_LIMIT)
		return fmt.Errorf("error limit exceeded %d", MAX_FAILURE_LIMIT)
	}

	log.Debug("Calling dummy_backoff method")
	to_return_error := false
	rand.Seed(time.Now().UnixNano())

	// Generate a random integer between 0 and 99
	randomInt := rand.Intn(100)
	if randomInt%5 != 0 {
		to_return_error = true
	}

	// Call the dummy_backoff method with the generated error flag
	success, _ := expobackoff.dummy_backoff(to_return_error)
	if success {
		log.Debug("Successful execution of dummy_backoff")
		return nil
	} else {
		log.Debug("Received error, starting backoff process")
		wait_time_sec := expobackoff.CurrentBackoffCount // Calculate wait time based on current backoff count
		expobackoff.CurrentBackoffCount++                // Increment the backoff count
		log.Debugf("Will wait for %d seconds before retrying", wait_time_sec)
		time.Sleep(time.Duration(wait_time_sec) * time.Second) // Wait for the calculated time
		return expobackoff.simple_backoff()                    // Retry the backoff process
	}
}

func (expobackoff *Expobackoff) wait_and_execute(time_to_wait int, to_return_error bool) {
	time.Sleep(time.Duration(time_to_wait) * time.Second)
	success, _ := expobackoff.dummy_backoff(to_return_error)
	if success {
		log.Debug("Successfull execution")
		expobackoff.TaskChan <- "Success"
	} else {
		expobackoff.TaskChan <- "Failure"
	}
}

func (expobackoff *Expobackoff) to_return_error() bool {
	rand.Seed(time.Now().UnixNano())

	// Generate a random integer between 0 and 99
	randomInt := rand.Intn(100)
	if randomInt%5 != 0 {
		return true
	}
	return false
}

func (expobackoff *Expobackoff) simple_backoff_async() error {
	log.Debug("Initiating function call")
	if expobackoff.CurrentBackoffCount > MAX_FAILURE_LIMIT {
		return fmt.Errorf("Error limit exceeded %d", MAX_FAILURE_LIMIT)
	}
	log.Debug("Call dummy method")

	to_return_error := expobackoff.to_return_error()
	go expobackoff.wait_and_execute(expobackoff.CurrentBackoffCount, to_return_error)
	success := <-expobackoff.TaskChan
	if success == "Success" {
		log.Debug("Successfull execution")
		return nil
	} else {
		log.Debug("recieved error starting backoff process")
		expobackoff.CurrentBackoffCount++
		return expobackoff.simple_backoff_async()
	}
}

func Test() {

	exp_back_off := Expobackoff{
		CurrentBackoffCount: 1,
		TaskChan:            make(chan string),
	}
	// err := exp_back_off.simple_backoff()
	// log.Info("Method call is:", err)
	err := exp_back_off.simple_backoff_async()
	log.Info("Method call is:", err)
}
