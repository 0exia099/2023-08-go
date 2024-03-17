package main

import "fmt"

func main() {
	c := Add(3, 6)
	fmt.Println(c)
	c, success := Divide(9, 3)
	fmt.Println(c, success)
	d, success := Divide(9, 0)
	fmt.Println(d, success)
}
func Add(a int, b int) int {
	return a + b
}
func Divide(a int, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return
	}
	result = a / b
	success = true
	return
}
