package main

import "fmt"

func main() {
	day := "tuesday"
	switch day { // break필요 없음. 밑에것도 실행시키고 싶으면 fallthrough사용
	case "monday", "tuesday":
		fmt.Println("월, 화요일은 수업 가는 날입니다.")
		fallthrough
	case "wednesday", "thursday":
		fmt.Println("수, 목요일은 실습 가는 날입니다.")
		fallthrough
	default:
		fmt.Println("aaa")
	}
}
