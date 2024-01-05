package main

import "fmt"

func main() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue) // close channel, can still print out below what is in it

	for elem := range queue {
		fmt.Println(elem)
	}
}
