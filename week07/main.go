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
	in := bufio.NewReader(os.Stdin)
	fmt.Print("input your score: ")
	i, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i)                //줄바꿈 등 제거 파이썬 strip 함수와 비슷
	score, _ := strconv.ParseInt(i, 16, 32) //16진 정수형 32bit
	if score >= 90 {
		fmt.Println("A")
		fmt.Printf("%d\n", score)
	} else {
		fmt.Println("BCDF")
		fmt.Printf("%d\n", score)
	}
}
