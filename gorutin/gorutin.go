package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func PrintHangul() {
	hanguls := []rune{'가', '나', '다', '라', '마', '바', '사'}
	for _, v := range hanguls {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c ", v)
	}
}
func PrintNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func SumAtoB(a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Printf("%d부터 %d까지 합계는 %d입니다.\n", a, b, sum)
	wg.Done()
}

func main() {
	// go PrintHangul()
	// go PrintNumbers()
	// time.Sleep(3 * time.Second)

	// for i := 0; i < 10; i++ {
	// 	SumAtoB(1, 1000000000)
	// }
	// fmt.Println("모든 계산이 완료되었습니다.1")

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go SumAtoB(1, 1000000000)
	}
	//time.Sleep(1 * time.Second)
	wg.Wait()
	fmt.Println("모든 계산이 완료되었습니다.2")
}
