package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	answer := rand.Intn(6) + 1 //1~6 random

	var win bool = false
	for guesses := 0; guesses < 3; guesses++ {
		fmt.Printf("%d번의 기회가 남음, 숫자 입력: ", 3-guesses)
		in := bufio.NewReader(os.Stdin)
		i, err := in.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		i = strings.TrimSpace(i)
		guess, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}

		if answer == guess {
			win = true
			fmt.Println("정답입니다!")
			break
		} else if answer > guess {
			fmt.Println("입력하신 수는 정답보다 작은 수 입니다. LOW")
		} else if answer < guess {
			fmt.Println("입력하신 수는 정답보다 큰 수 입니다. HIGH")
		}
	}
	if win == false {
		fmt.Printf("너가 짐. 정답은 %d임", answer)
	} else if win == true {
		fmt.Println("너가 이김")
	}

}
