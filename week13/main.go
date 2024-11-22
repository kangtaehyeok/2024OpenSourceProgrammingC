package main

import (
	"fmt"
	"os"
	"reflect"
)

// func test(strs string) {
func test(strs ...string) { //func test(i int, strs ...string)
	fmt.Println(strs, reflect.TypeOf(strs))
}

func main() {
	//fmt.Println(os.Args[1:], len(os.Args))
	slices := os.Args[1:]
	fmt.Println(slices[1])
	for _, slice := range slices {
		fmt.Println(slice)
	}
	slices = append(slices, "forever", "!")
	fmt.Println(slices, len(slices))
	test("ABC")
	test("ABC", "123")
	test()
	test("ABC", "123", "REAR")

}
