package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	os.Setenv("FOO", "1")
	fmt.Println("FOO: ", os.Getenv("FOO"))
	fmt.Println("BAR: ", os.Getenv("BAR"))

	fmt.Println()
	fmt.Println("os.Environ(): ", os.Environ())
	//os.Environ lists all key/value pairs in environment
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
