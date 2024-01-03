package main

//select cases aren't just read only and switch, they can do things in their case if it's legal
import "fmt"

// for this example the default will be hit on all three
func main() {
	messages := make(chan string)
	signals := make(chan bool)
	//if messages had a value, it would hit first case
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	//msg can't be sent into messages because it doesn't receive anything so falls to default
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	//neither messages or signals receive anything so this goes to default
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
