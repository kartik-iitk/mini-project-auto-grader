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
	mtx.Lock()
	pwd, _ := os.Getwd() // Get current working directory
	os.Chdir(path)       // Change it to the current path
	result, _ := exec.Command("/usr/local/go/bin/go", "test").Output()
	testResult := string(result)
	if testResult != "" {
		count := strings.Count(testResult, "FAIL")
		if count > 0 {
			count-- // One extra FAIL occurs in the string
		}
		count = 6 - count // There are a total of 6 test cases in task_test.go
		// We get the GitHub Username by trimming of the first part of the string.

		taskPassed[strings.TrimPrefix(d.Name(), "recruitment-task-")] = count

	}
	os.Chdir(pwd) // Reset the current directory for further walking
	mtx.Unlock()
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
					defer wg.Done()
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
