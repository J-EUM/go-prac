package main

import (
	"fmt"
	"strings"
)

func main() {
	courses := []string{
		"go Go 고고",
		"java",
		"python",
	}

	for _, x := range courses {
		// x 가 "Go" 를 포함하고 있으면 true
		if strings.Contains(x, "Go") {
			fmt.Println("Go is found in", x, "index is", strings.Index(x, "Go"))
		}
	}

	newString := "Go is a great programming language. Go for it!"

	// newString이 "Go"로 시작하면 true
	fmt.Println(strings.HasPrefix(newString, "Go"))
	// newString이 "Go"로 끝나면 true
	fmt.Println(strings.HasSuffix(newString, "Go")) // false
	fmt.Println(strings.HasSuffix(newString, "!"))  // true
	// newString에 "Go"가 몇개 있는지 숫자 반환
	fmt.Println(strings.Count(newString, "Go")) // 2
	// newString에서 첫 "Go"의 위치(인덱스), 없다면 -1
	fmt.Println(strings.Index(newString, "Go")) //0
	fmt.Println(strings.Index(newString, "oG")) // -1
	// newString에서 마지막 "Go"의 인덱스
	fmt.Println(strings.LastIndex(newString, "Go")) // 36

	// replace
	if strings.Contains(newString, "Go") {
		// 아래 두개 같은거
		// newString = strings.Replace(newString, "Go", "Golang", -1) // -1: 전부바꿔라, 1: 앞에서부터 한개만 바꿔라
		newString = strings.ReplaceAll(newString, "Go", "Golang")
	}

	fmt.Println(newString)

	// 스트링 비교 == != > <,...로 가능: 스트링은 rune(int32) 슬라이스기때문
	a := "A"
	if a == "A" {
		fmt.Println("a는 A")
	}

	// white space 제거
	badEmail := "  me@here.com    "
	badEmail = strings.TrimSpace(badEmail)
	fmt.Printf("=%s=\n", badEmail)

}
