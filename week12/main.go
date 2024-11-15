package main

import (
	"fmt"
	"time"
)

func main() {
	var scores [3]int
	scores[1] = 90
	fmt.Println(scores[1], scores[0], scores[2])
	fmt.Printf("%#v\n", scores)
	fmt.Println(scores)

	//var dates [3]time.Time
	//dates[1] = time.Unix(1918171615, 0)
	//fmt.Println(dates[1])

	//	dates := [3]time.Time{
	//		time.Unix(0, 0),
	//		time.Unix(1, 0),
	//		time.Unix(1918171615, 0), //need , comma
	//	}
	//	for i := 0; i < 3; i++ {
	//		fmt.Println(dates[i])
	//	}
	//}

	dates := [3]time.Time{
		time.Unix(0, 0),
		time.Unix(1, 0),
		time.Unix(1918171615, 0)}
	for i := 0; i < 3; i++ {
		fmt.Println(dates[i])
	}
	fmt.Printf("%#v\n", dates)
	fmt.Println(dates)
}
