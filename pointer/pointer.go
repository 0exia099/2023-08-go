package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

func ChangeData(arg *Data) {
	arg.value += 100
	arg.data[100] += 100
}

func main() {
	var a int = 500
	var p *int
	p = &a
	fmt.Printf("p의 값: %p\n", p)
	fmt.Printf("p가 가리키는 메모리의 값: %d\n", *p)
	*p = 100
	fmt.Printf("a의 값: %d\n", a)
	fmt.Printf("p가 가리키는 메모리의 값: %d\n", *p)

	var data Data
	ChangeData(&data)
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100]%d\n", data.data[100])
	// 함수에서 주소를 리턴하면 함수가 종료되면서 해당 주소가 스택에서 해제될 수 있음 -> c에서는 지원하지 않지만 golang에서는 해당 부분을 스택대신
	// 힙에 할당하여 가능
}
