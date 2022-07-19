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
	cmd := exec.Command("/usr/local/go/bin/go", "test")
	cmd.Dir = path // Set Directory for running the command.
	result, _ := cmd.Output()
	testResult := string(result)
	if testResult != "" {
		count := strings.Count(testResult, "FAIL")
		if count > 0 {
			count-- // One extra FAIL occurs in the string
		}
		count = 6 - count // There are a total of 6 test cases in task_test.go
		// We get the GitHub Username by trimming of the first part of the string.
		mtx.Lock() // As we will access the common map.
		taskPassed[strings.TrimPrefix(d.Name(), "recruitment-task-")] = count
		mtx.Unlock()
	}
}

func main() {
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
