package main

import "fmt"

var g int = 10 // 전역변수

func main() {
	var m = 20 // 지역변수 int형
	{
		var s int = 50 //지역변수
		fmt.Println(m, s, g)
	}
	// m = s+20 // Error s는 위의 중괄호 닫히면서 사라짐
}
