package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {

	// var i int = 55

	var f float32 = 3.99 //4바이트
	//float64 << 8바이트 double

	//var i int
	//i = 55

	i := 55

	fmt.Println(f, math.Ceil(3.49))
	fmt.Println(strings.Title("kim inha"))
	fmt.Printf("i는 %d\n", i)
	fmt.Print("i는 ", i, "\n")
	fmt.Println("i는", i)
}
