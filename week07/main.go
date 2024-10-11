package main

import (
	"fmt"
	"time"
)

func main() {
	// 오늘은 10월 11일
	var now time.Time = time.Now()
	fmt.Printf("%d년 %d월 %d일\n", now.Year(), now.Month(), now.Day())
	fmt.Printf("%d년 %d월 %d일 %d시 %d분 %d초임\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
}
