package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ { // 조건을 안적으면 무한반복
		fmt.Print(i, ", ")
	}
	fmt.Println()
	// golang은 while없음
	a := 1
	b := 1
Outerfor:
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				break Outerfor
			}
		}
	}
	fmt.Printf("%d * %d = %d", a, b, a*b)
	fmt.Println()
	for i := 1; i <= 7; i++ {
		for j := 7; j >= i; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
