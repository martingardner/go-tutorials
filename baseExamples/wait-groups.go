package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("worker %d startin\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

/*
WaitGroup has 3 methods
.Add() | .Wait() | .Done()
starts at 0, add adds to it, done subtracts from it, wait blocks any remaining processes until back to 0
*/

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		i := i

		go func() {
			//defer blocks done until worker finishes firing
			defer wg.Done()
			worker(i)
		}()
	}
	// will wait until the initial 5
	wg.Wait()
	fmt.Println("after wg.Wait()")
}
