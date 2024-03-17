package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	var a int
	var b int
	fmt.Print("정수를 두개 입력하세요 : ")
	n, err := fmt.Scan(&a, &b)
	if err != nil {
		stdin.ReadString('\n') // 버퍼 비우기
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}
