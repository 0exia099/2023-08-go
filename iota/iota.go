package main

import "fmt"

const (
	Red   int = iota // 0
	Grean int = iota // 1
	Blue  int = iota // 2
)
const (
	R = iota + 1 // 1
	G            // 2
	B            // 3
)

func main() {
	fmt.Println(Red, Grean, Blue)
	fmt.Println(R, G, B)
}
