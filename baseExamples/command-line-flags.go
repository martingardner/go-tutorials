package main

import (
	"flag"
	"fmt"
)

// go build command-line-flags.go  first
// -h default flag always on to provide help command, list of all arguments
func main() {

	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	var svar string
	//attaches a pointer to svar below to represent flags
	flag.StringVar(&svar, "svar", "bar", "a string var")
	//trigger this after attaching the flags to instantiate it
	flag.Parse()

	//without asterik, gives memory position
	fmt.Println("word without asterik: ", wordPtr)
	//need asterik to dereference and get the actual result
	fmt.Println("word: ", *wordPtr)
	fmt.Println("numb: ", *numbPtr)
	fmt.Println("fork: ", *forkPtr)
	fmt.Println("svar: ", svar)
	//tail will capture any additional comments that are on the command line that the flags aren't capturing.
	fmt.Println("tail: ", flag.Args())
}
