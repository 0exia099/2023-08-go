package main

import "fmt"

type House struct {
	Address string
	Size    int
	Price   float64
	Type    string
}
type AAA struct {
	Home House
	name string
}

func main() {
	var house House
	house.Address = "주소"
	house.Size = 28
	house.Price = 9.8
	house.Type = "아파트"
	fmt.Println("주소:", house.Address)
	fmt.Printf("크기: %d평\n", house.Size)
	fmt.Printf("가격: %2.f억원\n", house.Price)
	fmt.Println("타입:", house.Type)
	a := AAA{
		House{"주소1", 2, 3.3, "4"},
		"aaa",
	}
	fmt.Println(a.Home.Address, a.Home.Size, a.Home.Price, a.Home.Type, a.name)
	// 복사도 가능
}
