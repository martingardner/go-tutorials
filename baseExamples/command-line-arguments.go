package main

import (
	"fmt"
	"os"
)

// run go build file name first, then run that .exe   ./command-line-arguments a b c d
func main() {

	//first argument of argsWithProg is path to the program
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println("argsWithProg", argsWithProg)
	fmt.Println("argsWithoutProg", argsWithoutProg)
	fmt.Println("arg", arg)
}
