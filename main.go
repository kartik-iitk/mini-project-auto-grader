package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

var wg sync.WaitGroup
var mtx sync.Mutex
var taskStatus map[string]string

func findPassed(path string, userID string) {
	cmd := exec.Command("/usr/local/go/bin/go", "test", "./...") // "./... runs tests of nested directories!
	cmd.Dir = path                                               // Set Directory for running the command.
	result, err := cmd.Output()
	mtx.Lock()
	if err == nil {
		// You should never compare errors other than nil with == operator. Use errors.Is() function!
		taskStatus[userID] = "All Passed"
	} else {
		// In case, there is failure, instead of parsing the result for "FAIL" it is a better strategy to simply
		// log the result for manual inspection and flag as failed. This keeps the code reusable.
		log.Println(userID, ":", string(result))
		// Get error code. Avoid parsing using err.Errors() method.
		code := err.(*exec.ExitError).ExitCode() // type assertion coupled with ProcessorState function to get ErrCode.
		switch code {
		case 1:
			taskStatus[userID] = "Some Tests Failed"
		default:
			taskStatus[userID] = "Other Error"
		}
	}
	mtx.Unlock()
}

func main() {
	// Apply anti-cheating measure by executing the bash-script, which deletes all test files and puts in our test file.
	cmd := exec.Command("bash", "applyAntiCheatingMeasure.sh")
	cmd.Dir = "./anti-cheating"
	cmd.Run()

	// Initialisation
	taskStatus = make(map[string]string)

	// Begin testing
	maxDepth := 1 // For WalkDir function, 0 is the root directory from where we start executing the function.
	var err = filepath.WalkDir("./submission-data/",
		func(path string, d os.DirEntry, err1 error) error {
			if err1 != nil {
				return err1
			}
			if d.IsDir() == true && strings.Count(path, string(os.PathSeparator)) <= maxDepth {
				// We get the GitHub Username by trimming of the first part of the string.
				userID := strings.TrimPrefix(d.Name(), "recruitment-task-")
				wg.Add(1)
				go func() {
					// Wrapper function to implement wait-groups.
					defer wg.Done() // To handle exceptions better.
					findPassed(path, userID)
				}()
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	wg.Wait()
	fmt.Println(taskStatus)
}
