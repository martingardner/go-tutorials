package main

import (
	"fmt"
	"time"
)

func main() {
	//makes a request channel
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		//<-limiter knows that it is being sent into the requests channel
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	//makes a bursty limiter channel

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
		fmt.Println("burstyLimiter", i, time.Now())
	}

	go func() {
		for t := range time.Tick(800 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	//make burstyRequests channel
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

}
