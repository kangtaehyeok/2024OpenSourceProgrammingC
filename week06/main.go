package main

import (
	"fmt"
)

func main() {
	var f float64
	var i int
	var b bool
	var s string

	fmt.Println(f, i, b, s)
	fmt.Printf("%f %d %t %s\n", f, i, b, s) //zero value
	f = 2.7
	i = 3
	fmt.Print("\n", f < float64(i), "\n")
}
