package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

var taskStatus map[string]string
var pollingTime time.Duration

type pair struct {
	id  string
	res string
}

func findPassed(path string, userID string, c chan pair) {
	cmd := exec.Command("/usr/local/go/bin/go", "test", "./...") // "./... runs tests of nested directories!
	cmd.Dir = path                                               // Set Directory for running the command.
	result, err := cmd.Output()
	if err == nil {
		// You should never compare errors other than nil with == operator. Use errors.Is() function!
		c <- pair{userID, "All Passed"}
	} else {
		// In case, there is failure, instead of parsing the result for "FAIL" it is a better strategy to simply
		// log the result for manual inspection and flag as failed. This keeps the code reusable.
		log.Println(userID, ":", string(result))
		// Get error code. Avoid parsing using err.Errors() method.
		code := err.(*exec.ExitError).ExitCode() // type assertion coupled with ProcessorState function to get ErrCode.
		switch code {
		case 1:
			c <- pair{userID, "Some Tests Failed"}
		default:
			c <- pair{userID, "Other Error"}
		}
	}
}

func updateMap(c chan pair, exit chan int) {
	fmt.Println("Start updateMap", c)
	for {
		time.Sleep(pollingTime) // Channel Polling Frequency to reduce CPU usage.
		select {
		case p := <-c:
			taskStatus[p.id] = p.res
		case <-exit: // Something has been written on exit channel
			close(c)
			close(exit)
			return
		}
	}
}

func main() {
	// Apply anti-cheating measure by executing the bash-script, which deletes all test files and puts in our test file.
	cmd := exec.Command("bash", "applyAntiCheatingMeasure.sh")
	cmd.Dir = "./anti-cheating"
	cmd.Run()

	// Initialisation
	taskStatus = make(map[string]string)
	// Allows easy communication between goroutines, so that the map can be accessed while some goroutines execute!
	c := make(chan pair)
	exit := make(chan int)
	pollingTime = 1e6
	concurrency := 10
	// A buffered channel (semaphore) which can hold only 'concurrency' ints. Helps limit number of Goroutines.
	// Mutex limits things to one thread, while semaphores limits things to 'concurrency' number of threads.
	sem := make(chan int, concurrency)

	// Begin testing
	go updateMap(c, exit)
	maxDepth := 1 // For WalkDir function, 0 is the root directory from where we start executing the function.
	var err = filepath.WalkDir("./submission-data/",
		func(path string, d os.DirEntry, err1 error) error {
			if err1 != nil {
				return err1
			}
			if d.IsDir() == true && strings.Count(path, string(os.PathSeparator)) <= maxDepth {
				// We get the GitHub Username by trimming of the first part of the string.
				userID := strings.TrimPrefix(d.Name(), "recruitment-task-")
				sem <- 1 // Same as wg.Add(), we add a new int to the channel.
				go func() {
					// Wrapper function to implement the buffered channels.
					// We could have simply written "<- sem" after findPassed call.
					// But we use deferred calls to handle exceptions better.
					defer func() { <-sem }() // To remove one int from buffered channel.
					findPassed(path, userID, c)
				}()
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	for i := 0; i < cap(sem); i++ {
		// When the last goroutine is pushed to the buffered channel, there are still 'concurrency' amount of goroutines
		// running so we need to wait for them to finish. Similar to wg.Wait() function.
		sem <- 1
	}
	time.Sleep(pollingTime) // Prevent main() from terminating before updateMap() can update the map for the last entry.
	exit <- 0               // Exit from updateMap() by pushing 0 to the channel
	fmt.Println(taskStatus)
}
