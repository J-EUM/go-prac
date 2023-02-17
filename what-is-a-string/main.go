package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println()
	name := "Hello world"
	fmt.Println("String: ", name)
	fmt.Println()

	// string: a slice of bytes
	fmt.Println("Bytes")
	for i := 0; i < len(name); i++ {
		fmt.Printf("%x ", name[i])
	}

	fmt.Println()
	fmt.Println()

	// string: a slice of runes.
	// rune: 32bit integer 룬 유니코드값 저장가능
	fmt.Println("Index\tRune\tSting")
	for x, y := range name {
		fmt.Println(x, "\t", y, "\t", string(y))
	}
	fmt.Println()

	name = "안녕하세요요"
	fmt.Println("Index\tRune\tSting")
	for x, y := range name {
		fmt.Println(x, "\t", y, "\t", string(y))
	}

	fmt.Println()
	fmt.Println()

	fmt.Println("문자열붙이기 3방법")

	h := "Hello, "
	w := "world."

	// 스트링은 immutable -> 별로 좋은방법아니다
	myString := h + w
	fmt.Println(myString)

	// using fmt - +보다 효율적
	myString = fmt.Sprintf("%s%s", h, w)
	fmt.Println(myString)

	// using stringbuilder - 매우 efficient, more control
	var sb strings.Builder
	sb.WriteString(h)
	sb.WriteString(w)
	fmt.Println(sb.String())

	fmt.Println()

	// substring
	name = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	fmt.Println("Getting a substring")
	fmt.Println(name[10:13])

	fmt.Println()
	fmt.Println()

	// 인덱스
	//             01234567890123456789012345678901234
	courseName := "Learn Go for Beginners Crash Course"

	for i := 13; i <= 21; i++ {
		// string() 캐스트를 하지 않으면 룬(int32)형태로 출력
		fmt.Print(string(courseName[i]))
	}

}
