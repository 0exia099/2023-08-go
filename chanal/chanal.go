package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// var wg sync.WaitGroup
	// ch := make(chan int) // 채널 생성(크기 0인 채널(버퍼 없음))
	// // ch := make(chan int, 2)	// 크기 2인 채널(버퍼가짐. 2개의 메시지 보관 가능)
	// // 버퍼가 모두 차면 빈자리 생길때 까지 기다림
	// wg.Add(1)
	// go square(&wg, ch) // go루틴 생성
	// ch <- 9            // 채널에 데이터 추가
	// // 채널에 넣은 후, 해당 데이터를 다른 고루틴에서 가져갈때 까지 블로킹상태로 대기
	// // 버퍼가 있는 채널이면 버퍼모두 차기전까지 실행된다
	// wg.Wait()          // 끝나기 기다림(없으면 square가 다 실행되기 전에 종료됨)
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	close(ch) // 채널 닫음
	wg.Wait()
}
func square(wg *sync.WaitGroup, ch chan int) {
	// n := <-ch // 채널에서 데이터 빼옴(채널이 비어있으면 데이터 들어올때까지 블라킹된다)

	// time.Sleep(time.Second) // 1초 기다림
	// fmt.Printf("Square: %d\n", n*n)
	// wg.Done()

	for n := range ch { // 채널이 닫히기 전까지 반복된다
		fmt.Printf("Square: %d\n", n*n)
		time.Sleep(time.Second)
	}
	wg.Done()
}
