package main

import "fmt"

func receiveNums(ch chan int) {
	for i := 0; i <= 10; i++ {
		ch <- i
		if i == 10 {
			close(ch)
		}
	}
}

func printNums(ch chan int, done chan bool) {
	for elem := range ch {
		fmt.Println(elem) // print the received numbers from the unbuffered channel
	}
	done <- true // only send `true` once all numbers from the `ch` channel have finished printing
}

func main() {
	ch := make(chan int)    //unbuffered channel declaration
	done := make(chan bool) //unbuffered channel used for signalling completion

	go receiveNums(ch)

	go printNums(ch, done)

	//only close the done channel once the `printNums` funciton has finished its execution,
	//which would effectively force the main program to wait until that happens
	<-done
}
