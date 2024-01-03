package main

import "fmt"

//first param only accepts a channel that sends a value
func ping(pSend chan<- string, msg string){
	pSend <- msg
}

func pong(pGet <-chan string, pSend chan<- string){
	msg := <-pGet //channel for receive
	pSend <- msg //channel for send
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}