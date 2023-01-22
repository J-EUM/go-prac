package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

func main() {
	reader = bufio.NewReader(os.Stdin)

	// fmt.Println("What is your name?")
	// prompt() // fmt.Print("-> ")

	// userInput, _ := reader.ReadString('\n')
	// userInput = strings.Replace(userInput, "\n", "", -1)
	userName := readString("What is your name?")

	age := readInt("How old are you?")

	// +로 연결 concatination하면 느리고 메모리도 많이쓴다. ,로연결하면 사이사이 스페이스들어간다
	// fmt.Println("Your name is "+userName+". You are", age, "years old.")

	// Sprintf: S스트링 프린트 f포맷
	//fmt.Println(fmt.Sprintf("Your name is %s. You are %d years old.", userName, age))
	// fmt.Println(fmt.Sprintf()) == fmt.Printf()
	fmt.Printf("Your name is %s. You are %d years old.\n", userName, age)
}

func prompt() {
	fmt.Print("-> ")
}

func readString(s string) string {
	for {

		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "" {
			fmt.Println("똑바로치셈")
		} else {
			return userInput
		}
	}
}

func readInt(s string) int {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)

		// 유저인풋(스트링)을 int로 변환
		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Please enter a whole number")
		} else {
			return num
		}
	}
}
