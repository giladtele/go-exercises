package main

import "fmt"

func RoutinePrint(done chan bool) {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	done <- true // Send a signal after finishing the loop.
}

func main() {
	done := make(chan bool)
	go RoutinePrint(done)
	<-done // Wait for the signal before exiting the main function.
}
