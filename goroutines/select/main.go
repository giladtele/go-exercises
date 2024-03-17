package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func sendNums(numsChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()          // `defer` makes sure `wg.Done()` will only execute when the function nears its return phase
	for i := 0; i < 4; i++ { // Loop 4 times to send 4 random numbers
		numsChan <- rand.Intn(100)
	}
}

func sendStr(strChan chan string, wg *sync.WaitGroup) {
	defer wg.Done() // `defer` makes sure `wg.Done()` will only execute when the function nears its return phase
	strChan <- "hello"
}

func selectChans(numsChan chan int, strChan chan string) {
	for {
		select {
		case num, ok := <-numsChan:
			if !ok {
				numsChan = nil
			} else {
				fmt.Println(num)
			}
		case str, ok := <-strChan:
			if !ok {
				strChan = nil
			} else {
				fmt.Println(str)
			}
		}
		if numsChan == nil && strChan == nil {
			break
		}
	}
}

func main() {
	startTime := time.Now()

	rand.Seed(time.Now().UnixNano()) // Ensure different random numbers each run
	var wgNums, wgStr sync.WaitGroup
	numsChan := make(chan int)
	strChan := make(chan string)

	go selectChans(numsChan, strChan)

	for i := 0; i < 5; i++ { // Start 5 goroutines that trigger `sendNums` each
		wgNums.Add(1)
		go sendNums(numsChan, &wgNums)
	}

	wgStr.Add(1)
	go sendStr(strChan, &wgStr)

	wgNums.Wait()
	close(numsChan)

	endTime := time.Now()

	executionTime := endTime.Sub(startTime)
	fmt.Printf("Program executed in%v\n", executionTime)
}
