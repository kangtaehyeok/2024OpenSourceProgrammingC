package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("input number: ")
	in := bufio.NewReader(os.Stdin)
	i, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i) //줄바꿈 등 제거 파이썬 strip 함수와 비슷
	n, err := strconv.Atoi(i)

	if err != nil {
		log.Fatal(err)
	}

	var isPrime bool = true

	if n <= 1 {
		isPrime = false
	} else {
		j := 2
		for j < n {
			if n%j == 0 {
				isPrime = false
				fmt.Printf("%d ", j) // 약수
				//break performance up
			}
			j++
		}
	}

	if isPrime {
		fmt.Printf("\n%d is prime number.", n)
	} else {
		fmt.Printf("\n%d is NOT prime number", n)
	}
}
