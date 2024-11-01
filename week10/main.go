package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isPrime(n int) bool {
	if n <= 1 {
		//isPrime = false
		return false
	} else if n == 2 {
		//isPrime = true
		return true
	} else if n%2 == 0 {
		//isPrime = false
		return false
	} else {
		j := 3
		for j*j <= n {
			if n%j == 0 {
				//isPrime = false
				//break
				return false
			}
			fmt.Printf("%d ", j)
			j = j + 2
		}
	}
	return true
}

func main() {

	//fmt.Printf("%f\n", math.Sqrt(16.0))
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

	if isPrime(n) {
		fmt.Printf("\n%d is prime number.", n)
	} else {
		fmt.Printf("\n%d is NOT prime number", n)
	}
}
