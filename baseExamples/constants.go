package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 5000
	// works like js const, if you try to revalue, you get an error
	const d = 3e20 / n
	fmt.Println(d)
	
	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}