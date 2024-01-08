package main

//this example from Go has issues, if you don't comment out the panic on 24 then it breaks.
import (
	"fmt"
	"os"
)

/*
Immediately after getting a file object with createFile, we defer the closing of that file with closeFile.
This will be executed at the end of the enclosing function (main), after writeFile has finished.
*/
func main() {

	f := createFile("/tmp/defer.txt")
	defer closeFile(f) //defer basically pushes to end of list of things to do.  Would assume multi defer would go in sequential order of defers
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		//panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")

}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
