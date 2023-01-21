// 엔터입력을 받고
// 변수를 활용해서 응답을 출력
// 변수를 사용하는 이유: 프로그램 관리. 숫자들을 변수로 관리하지 않았으면 랜덤숫자로 바꿀때 겁나많이 바꿨어야함

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// 상수는 선언하고 사용하지 않아도 에러 안남
const prompt = "and don't type your number in, just press ENTER when ready."

func main() {
	// seed the random number generator
	// 이걸 하지 않으면 아래 firstNumber, secondNumber, subtraction은 항상 같은 숫자가 나온다.
	rand.Seed(time.Now().UnixNano())

	var firstNumber = rand.Intn(8) + 2 // rand.Intn(n)은 0~n의 숫자를 리턴한다. 게임에서 firstNumber를 곱할거니까 0이나 1을 곱하지 않게 하기 위해 firstNumber가 2~10이 되도록 이렇게~
	var secondNumber = rand.Intn(8) + 2
	var subtraction = rand.Intn(8) + 2
	var answer int // 값을 주지 않으면 기본값 0이 들어간다. string이면 빈스트링

	fmt.Println(firstNumber, secondNumber, subtraction)

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
