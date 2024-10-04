package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 13
	var f float64 = 12.9
	c1 := 'Z' //90
	c2 := '김' //44608 Unicode

	fmt.Println(c1, c2)

	fmt.Printf("value i : %d, value f : %f\n", i, f)
	// fmt.Printf("%d * %f = %f\n", i, f, i*f)
	fmt.Printf("float64: %d * %f = %f\n", i, f, float64(i)*f)
	fmt.Printf("int: %d * %f = %d\n", i, f, i*int(f))
	fmt.Println(reflect.TypeOf(i), reflect.TypeOf(f), reflect.TypeOf(c1), reflect.TypeOf(c2))
}
