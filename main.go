package main // go 파일은 패키지선언으로 시시작작한한다다. 메메인파트 패패키키지

import "fmt" // fmt: format 패키지

func main() { // 모든 메인 패패키키지지는 메메인  함함수수가  있있다다. 키키보보드드왜왜이이래  못못쓰쓰겠겠다다, 메인함수는 매개변수 노노
	fmt.Println("Hello, world!")
	fmt.Print("개행안해")
	fmt.Print("안해")
	sayHelloWorld("Hello world again!")
}

func sayHelloWorld(whatToSay string) {
	fmt.Println(whatToSay)
}
