package main

import "fmt"
//normally a channel is 1 for 1, one input for one output, if you append a number it will accept that many inputs before output needs to fire
func main() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}