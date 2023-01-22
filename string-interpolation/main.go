package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

var reader *bufio.Reader

// 사용자정의타입
type User struct {
	UserName       string
	Age            int
	FavoriteNumber float64
	// 사용하지 않아도 상관없음
	OwnsADog bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var user User

	// fmt.Println("What is your name?")
	// prompt() // fmt.Print("-> ")

	// userInput, _ := reader.ReadString('\n')
	// userInput = strings.Replace(userInput, "\n", "", -1)
	user.UserName = readString("What is your name?")
	user.Age = readInt("How old are you?")
	user.FavoriteNumber = readFloat("What is your favorite number?")
	user.OwnsADog = readBool("Do you own a dog (y/n)?")

	// +로 연결 concatination하면 느리고 메모리도 많이쓴다. ,로연결하면 사이사이 스페이스들어간다
	// fmt.Println("Your name is "+userName+". You are", age, "years old.")

	// Sprintf: S스트링 프린트 f포맷
	//fmt.Println(fmt.Sprintf("Your name is %s. You are %d years old.", userName, age))
	// fmt.Println(fmt.Sprintf()) == fmt.Printf()
	fmt.Printf("Your name is %s, and you are %d years old. Your favorite number is %.2f. Owns a dog: %t.\n",
		user.UserName,
		user.Age,
		user.FavoriteNumber,
		user.OwnsADog)
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

func readFloat(s string) float64 {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)

		// 유저인풋(스트링)을 int로 변환
		num, err := strconv.ParseFloat(userInput, 64)
		if err != nil {
			fmt.Println("Please enter a number")
		} else {
			return num
		}
	}
}

// 키보드입력 y면 true n면 false리턴
func readBool(s string) bool {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	for {
		fmt.Println(s)
		prompt()

		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if char == 'n' || char == 'N' {
			return false
		} else if char == 'y' || char == 'Y' {
			return true
		} else {
			fmt.Println("y/n만 입력하셈")
		}
	}

}
