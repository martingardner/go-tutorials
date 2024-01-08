package main

import (
	"fmt"
	"slices"
)

func main() {
	strs := []string{"c", "a", "b"}
	fmt.Println("strs before sort: ", strs)
	fmt.Println("strs IsSorted before: ", slices.IsSorted(strs))
	slices.Sort(strs)
	fmt.Println("Strings: ", strs)
	fmt.Println("--------")

	ints := []int{7, 2, 4}
	fmt.Println("ints before sort: ", ints)
	fmt.Println("ints IsSorted before: ", slices.IsSorted(ints))
	slices.Sort(ints)
	fmt.Println("Ints: ", ints)
	fmt.Println("--------")
	//IsSorted checks if an array has already been sorted
	s := slices.IsSorted(ints)
	fmt.Println("Sorted: ", s)
}
