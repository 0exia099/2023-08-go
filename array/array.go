package main

import "fmt"

func main() {
	var arr1 [5]int = [5]int{10, 20, 30, 40, 50}
	var arr2 = []int{10, 20, 30, 40, 50}
	arr3 := []int{10, 20, 30, 40, 50}
	var arr4 [5]int = [5]int{1: 10, 3: 20} // 원하는 인덱스만 초기화(나머지는 0)
	arr5 := [...]int{10, 20, 30, 40, 50}   // ...을 쓰면 초기화하는 원소의 갯수만큼으로 설정된다.
	for i := 0; i < 5; i++ {
		fmt.Print(arr1[i], " ")
	}
	fmt.Println()
	for i := 0; i < 5; i++ {
		fmt.Print(arr2[i], " ")
	}
	fmt.Println()
	for i := 0; i < 5; i++ {
		fmt.Print(arr3[i], " ")
	}
	fmt.Println()
	for i := 0; i < 5; i++ {
		fmt.Print(arr4[i], " ")
	}
	fmt.Println()
	for index, value := range arr5 { // range로 순회 가능
		fmt.Println(index, value)
	}
	for _, value := range arr5 { // _로 인덱스는 안받을 수 있다
		fmt.Println(value)
	}
	arr1 = arr4	// 배열크기 같으면 복사 가능
	
}
