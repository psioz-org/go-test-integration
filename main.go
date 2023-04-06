package main

import "fmt"

func main() {
	fmt.Printf("result: %v", sum(3, 4))
}

func sum(a int, b int) int {
	return a + b
}
