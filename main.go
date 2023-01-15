package main // go 파일은 패키지선언으로 시시작작한한다다. 메메인파트 패패키키지

import (
	"fmt" // fmt: format 패키지
	"myapp/doctor"
)

func main() {
	// var whatToSay string
	// whatToSay = doctor.Intro()
	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)
}
