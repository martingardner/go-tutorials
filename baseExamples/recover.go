package main

import "fmt"

func mayPanic() {
	panic("a problem")
}

func main() {
	//recover has to be in a defer function
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	fmt.Println("After mayPanic()")
}
