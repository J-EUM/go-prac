package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)

	// // 조건 없으면 무한반복
	// for {
	// 	fmt.Print("-> ")

	// 	// ReadString은 string, error를 반환한다. 지금은 error안쓸거임
	// 	userInput, _ := reader.ReadString('\n')

	// 	// 유저인풋에서 모든엔터를(-1) 공백으로 치환
	// 	userInput = strings.Replace(userInput, "\n", "", -1)

	// 	if userInput == "quit" {
	// 		break
	// 	} else {
	// 		fmt.Println(userInput)
	// 	}
	// }

	// 하나의 키입력을 받기위해 외부패키지 사용

	err := keyboard.Open()

	// 에러가 없으면 err값은 nil이다.
	// if 에러가 나면~
	if err != nil {
		log.Fatal(err)
	}

	// defer: 즉시 실행되지 않고 defer키워드가 있는 현재 함수(main)가 끝나고 나면 실행된다.
	// 익명함수
	// main함수가 끝나면 keyboard를 닫아라. 에러리턴 필요없고
	defer func() {
		_ = keyboard.Close()
	}()

	// map[인덱스타입]밸류타입
	coffees := make(map[int]string)
	coffees[1] = "Cappucino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Macchiato"
	coffees[6] = "Espresso"
	coffees[1] = "Cappucino"

	fmt.Println("MENU")
	fmt.Println("----")
	fmt.Println("1 - Cappucino")
	fmt.Println("2 - Latte")
	fmt.Println("3 - Americano")
	fmt.Println("4 - Mocha")
	fmt.Println("5 - Macchiato")
	fmt.Println("6 - Espresso")
	fmt.Println("Q - Quit the program")

	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if char == 'q' || char == 'Q' {
			break
		}

		// A알파벳을(string(char)) i인티저로 바꿔준다. string(char): rune타입 char를 스트링으로 바꿔줌
		i, _ := strconv.Atoi(string(char))

		// fmt.Sprintf는 스트링을 리턴한다.
		// char는 rune타입....이걸 string으로 변환해서 프린트하기. %q는 rune타입
		// t := fmt.Sprintf("You chose %q", char)
		// %d int타입
		// t := fmt.Sprintf("You chose %d", i)
		// %s string타입
		// t := fmt.Sprintf("You chose %s", coffees[i])

		// 강의에서는 이렇게 Println(fmt.Sprintf)를 하라는데.....vscode에서는 따로하란다
		// 나도 따로하는게 가독성이 좋지않나싶은데~~~
		fmt.Println(fmt.Sprintf("You chose %s", coffees[i]))
	}

	fmt.Println("Program exiting...")
}
