package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func square(ch chan int, quit chan bool) {
	for {
		select { // ch와 quit를 모두 기다림
		case n := <-ch: // ch는 받아서 출력
			fmt.Printf("Square: %d\n", n*n)
			time.Sleep(time.Second)
		case <-quit: // quit이면 종료
			wg.Done()
			return
		}
	}
}

func main() {
	ch := make(chan int)
	quit := make(chan bool)

	wg.Add(1)
	go square(ch, quit)
	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	quit <- true // 종료 채널 true넣음
	wg.Wait()
}
