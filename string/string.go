package main

import "fmt"

func main() {
	var char rune = '한' // rune : 단일 문자 저장하기위한 타입. int32와 동일
	fmt.Printf("%T\n", char)
	fmt.Println(char)
	fmt.Printf("%c\n", char)
}
