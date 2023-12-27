package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	var b,c int = 1,2
	fmt.Println(b,c)

	b = 3
	fmt.Println(b)
	//if you revalue d to something other than a boolean it will break
	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
}