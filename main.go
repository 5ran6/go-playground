package main

import (
	"fmt"
	"time"
)
//Implement the execute function that executes a func with max threads only - without wait group
func execute(funcAry []func(), maxThreads int) {
	semaphore := make(chan struct{}, maxThreads)

	done := make(chan struct{}, len(funcAry))


	for _, f := range funcAry {
		go func(f func()) {
			semaphore <- struct{}{}
			f()
			<-semaphore
			done <- struct{}{}
		}(f)
	}

	for i := 0; i < len(funcAry); i++ {
		<-done
	}
	close(semaphore)
}

func main() {
	// Create a slice of functions to test with delays and logging
	funcAry := []func(){
		func() {
			fmt.Println("Function 1 started")
			time.Sleep(5 * time.Second)
			fmt.Println("Function 1 completed")
		},
		func() {
			fmt.Println("Function 2 started")
			time.Sleep(4 * time.Second)
			fmt.Println("Function 2 completed")
		},
		func() {
			fmt.Println("Function 3 started")
			time.Sleep(1 * time.Second)
			fmt.Println("Function 3 completed")
		},
		func() {
			fmt.Println("Function 4 started")
			time.Sleep(2 * time.Second)
			fmt.Println("Function 4 completed")
		},
		func() {
			fmt.Println("Function 5 started")
			time.Sleep(3 * time.Second)
			fmt.Println("Function 5 completed")
		},
	}
	

	maxThreads := 3
	execute(funcAry, maxThreads)
}