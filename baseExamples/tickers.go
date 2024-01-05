package main

import (
	"fmt"
	"time"
)

func main() {
	//ticker is time per interval so it's akin to JS setInterval
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)
	// have to remember goroutines run separately so sleep is running at some time as this is doing it's tick interval
	go func() {
		for {
			//runs based on value not switch statement
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()
	// so when this sleep is over then it shuts down the goroutine
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	//done <- true not needed in this example but showing you're more than likely to add something to the channel before shutting it down.
	done <- true
	fmt.Println("Ticker stopped")
}
