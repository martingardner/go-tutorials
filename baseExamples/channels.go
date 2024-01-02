package main

import "fmt"

func main() {
	//make is the pretext for a channel
	messages := make(chan string)

	// <- sends a value into the messages channel
	go func() { messages <- "ping"}()

	// if it's <-channel then it receives a value from a channel
	msg := <-messages
	fmt.Println(msg)
}