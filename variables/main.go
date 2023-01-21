// 엔터입력을 받고
// 변수를 활용해서 응답을 출력

package main

import (
	"bufio"
	"fmt"
	"os"
)

// 상수는 선언하고 사용하지 않아도 에러 안남
const prompt = "and press ENTER when ready."

func main() {
	var firstNumber = 2
	var secondNumber = 5
	var subtraction = 7
	var answer int // 값을 주지 않으면 기본값 0이 들어간다. string이면 빈스트링

	reader := bufio.NewReader(os.Stdin)

	// 환영/설명
	fmt.Println("Guess the Number Game")
	fmt.Println("----------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	// 게임
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)
	reader.ReadString('\n')

	// 정답
	answer = firstNumber*secondNumber - subtraction
	fmt.Println("The answer is", answer)
}
