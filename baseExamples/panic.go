package main

import "os"

func main() {
	//panic is similar to js throw exception, everything after it is ignored, it just throws
	panic("a problem")
	//this is unreachable because panic automatically kills the program
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
