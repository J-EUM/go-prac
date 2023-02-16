package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	rand.Seed(time.Now().UnixNano())
	i := 1000
	for i > 100 {
		// 1~1001 랜덤숫자
		i = rand.Intn(1000) + 1
		//fmt.Println("i is", i)
		if i > 100 {
			fmt.Println("i is", i, "so loop keeps going")
		}
	}

	fmt.Println("i is", i, "빠이")
}
