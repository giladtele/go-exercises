package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func receiveNums(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()          // `defer` makes sure `wg.Done()` will only execute when the function nears its return phase
	for i := 0; i < 4; i++ { // Loop 4 times to send 4 numbers
		ch <- rand.Intn(100)
	}
}

func printNums(ch chan int) {
	for elem := range ch {
		fmt.Println(elem)
	}
}

func main() {
	startTime := time.Now()

	rand.Seed(time.Now().UnixNano()) // Ensure different random numbers each run
	var wgRandom sync.WaitGroup
	ch := make(chan int)

	go printNums(ch)

	for i := 0; i < 5; i++ { // Start 5 goroutines
		wgRandom.Add(1)
		go receiveNums(ch, &wgRandom)
	}

	wgRandom.Wait()
	close(ch)

	endTime := time.Now()

	executionTime := endTime.Sub(startTime)
	fmt.Printf("Program executed in%v\n", executionTime)
}
