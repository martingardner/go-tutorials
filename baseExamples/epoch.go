package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("now", now)

	fmt.Println("now.Unix()", now.Unix())
	fmt.Println("now.UnixMilli()", now.UnixMilli())
	fmt.Println("now.UnixNano()", now.UnixNano())

	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}
