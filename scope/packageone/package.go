package packageone

import "fmt"

// // 소문자로 시작: private, 이 패키지 안에서만 사용가능
// // 대문자로 시작: public, packageone을 임포트한 다른 패키지에서 사용가능

// var privateVar = "I am private"
// var PublicVar = "I am public (or exported)"

// func notExported() {

// }

// func Exported() {

// }

var PackageVar = "PackageVar"

func PrintMe(a, b, c string) {
	fmt.Println(a, b, c)
}
