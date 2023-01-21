package main

import "myapp/packageone"

// import (
// 	"fmt"
// 	"myapp/packageone"
// )

// // 패키지레벨변수
// var one = "One"

// func main() {
// 	// 블록레벨
// 	var one = "this is a block level variable"
// 	fmt.Println(one)
// 	myFunc()

// 	newString := packageone.PublicVar
// 	fmt.Println("From packageone:", newString)

// 	packageone.Exported()

// }

// func myFunc() {
// 	//var one = "the number one"
// 	fmt.Println(one)
// }

var myVar = "myVar"

// challenge
func main() {
	// variables
	// declare a package level variable for the main package named myVar

	// declare a block level variable for the main function called blockVar
	var blockVar = "blockVar"

	// declare a package level variable in the packageone package named PackageVar

	// create an exported function in packageone called PrintMe

	// in the main function, print out the values of myVar, blockVar, and PackageVar on one line using the PrintMe function in packageone
	packageone.PrintMe(myVar, blockVar, packageone.PackageVar)
}
