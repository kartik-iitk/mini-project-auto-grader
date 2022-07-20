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
var taskPassed map[string]int

func findPassed(path string, d os.DirEntry) {
	userID := strings.TrimPrefix(d.Name(), "recruitment-task-")
	if userID == d.Name() {
		// Ignore all the other test case files present in nested directories as the directory name will not start
		// recruitment-task- prefix!
		return
	}
	cmd := exec.Command("/usr/local/go/bin/go", "test")
	cmd.Dir = path // Set Directory for running the command.
	result, err := cmd.Output()
	if werr, ok := err.(*exec.ExitError); ok {
		// To handle the case where things other than task failed/succeeded happened. Eg: Build failed, etc.
		if s := werr.Error(); s != "exit status 0" && s != "exit status 1" {
			log.Println(path, ":", string(result)) // Print the log and ignore this for manual inspection.
			return
		}
	}
	testResult := string(result)
	if testResult != "" {
		count := strings.Count(testResult, "FAIL")
		if count > 0 {
			count-- // One extra FAIL occurs in the string
		}
		// There are a total of 6 test cases in task_test.go, we assume there is no other test
		// file (as the bash script ensures this).
		count = 6 - count
		// We get the GitHub Username by trimming of the first part of the string.
		mtx.Lock() // As we will access the common map.
		taskPassed[userID] = count
		mtx.Unlock()
	}
}

func main() {
	// Apply anti-cheating measure by executing the bash-script, which deletes all test files and puts in our test file.
	cmd := exec.Command("bash", "applyAntiCheatingMeasure.sh")
	cmd.Dir = "./anti-cheating"
	cmd.Run()

	// Begin testing
	taskPassed = make(map[string]int)
	var err = filepath.WalkDir("./submission-data",
		func(path string, d os.DirEntry, err1 error) error {
			if err1 != nil {
				return err1
			}
			if d.IsDir() == true {
				wg.Add(1)
				go func() {
					// Wrapper function to implement waitgroups.
					defer wg.Done() // To handle exceptions better.
					findPassed(path, d)
				}()
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	wg.Wait()
	fmt.Println(taskPassed)
}
