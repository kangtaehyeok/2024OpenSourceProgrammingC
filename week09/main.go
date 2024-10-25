package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//fmt.Sprintf("%f 나누기 %f는 %f입니다", 1.0, 3.0, 1.0/3.0) //console 출력
	result := fmt.Sprintf("%0.1f 나누기 %0.1f는 %0.1f입니다\n", 1.0, 3.0, 1.0/3.0) //서식을 문자열로 리턴
	fmt.Print(result)

	i := 1
	for i <= 10 {
		//fmt.Printf("%2d\n", i)
		fmt.Printf("%4d\n", rand.Intn(1000)+1)
		i++
	}

}
