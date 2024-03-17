package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3}
	for i := 0; i < len(slice); i++ {
		slice[i] += 10
	}
	for i, v := range slice {
		slice[i] = v * 2
		fmt.Println(slice[i])
	}
	fmt.Println(slice, len(slice))
	slice = append(slice, 4)
	fmt.Println(slice, len(slice))
	slice = append(slice, 5, 6, 7, 8, 9)
	fmt.Println(slice, len(slice))

	var slice1 = make([]int, 3, 5)
	slice2 := append(slice1, 4)
	slice1[0] = 100
	fmt.Println(slice1, len(slice1), cap(slice1))
	fmt.Println(slice2, len(slice2), cap(slice2))

	slice1 = make([]int, 3, 3)
	slice2 = append(slice1, 4)
	slice1[0] = 100
	fmt.Println(slice1, len(slice1), cap(slice1))
	fmt.Println(slice2, len(slice2), cap(slice2))
}
