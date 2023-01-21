package main

import (
	"fmt"
	"myapp/packageone"
)

// 패키지레벨변수
var one = "One"

func main() {
	// 블록레벨
	var one = "this is a block level variable"
	fmt.Println(one)
	myFunc()

	newString := packageone.PublicVar
	fmt.Println("From packageone:", newString)

	packageone.Exported()

}

func myFunc() {
	//var one = "the number one"
	fmt.Println(one)
}
