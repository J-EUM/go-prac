package main // go 파일은 패키지선언으로 시시작작한한다다. 메메인파트 패패키키지

import "fmt" // fmt: format 패키지

func main() {
	// var whatToSay string
	// whatToSay = "Hello world, again!"

	whatToSay := "Hello world, again!" // 알아서 스트링 할당

	sayHelloWorld(whatToSay)
}

func sayHelloWorld(whatToSay string) {
	fmt.Println(whatToSay)
}
