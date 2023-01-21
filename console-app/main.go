package main

import (
	"fmt"
	"log"

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

	fmt.Println("Press any key on the keyboard. Press ESC to quit.")

	for {
		char, key, err := keyboard.GetSingleKey()

		if err != nil {
			log.Fatal(err)
		}

		if key != 0 {
			fmt.Println("You pressed", char, key)

		} else {
			fmt.Println("You pressed", char)
		}

		if key == keyboard.KeyEsc {
			break
		}
	}

	fmt.Println("Program exiting...")
}
