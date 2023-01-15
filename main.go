package main // go 파일은 패키지선언으로 시시작작한한다다. 메메인파트 패패키키지

import (
	"bufio"
	"fmt" // fmt: format 패키지
	"myapp/doctor"
	"os"
	"strings"
)

func main() { // 메인패키지는 메인함수를 가져야 한다
	reader := bufio.NewReader(os.Stdin) // 입력을 받는다 ()어디를 읽을지 알려줘야함

	// var whatToSay string
	// whatToSay = doctor.Intro()
	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)

	// go는 반복문 for밖에 없음
	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n') // 유저가 타이핑하는거 캡쳐. 타이핑이 언제 끝나냐: 엔터. ''싱글쿼트 써야함

		userInput = strings.Replace(userInput, "\n", "", -1) // 개행 발견할때마다(-1) 공백으로 치환

		if userInput == "quit" {
			break
		} else {
			// response := doctor.Response(userInput)
			// fmt.Println(response)
			fmt.Println(doctor.Response(userInput))
		}
	}
}
