package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	//makes a channel
	done := make(chan bool, 1)
	// starts a goroutine
	go worker(done)
	// this blocks ending the program until done is hit in worker, otherwise main would end before worker finished doing what it was doing.
	<-done
}