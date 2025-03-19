package main

import "fmt"

func main() {
	for i, j := 0, 5; i < j; i, j = i+1, j-1 {
		fmt.Printf("i=%d, j=%d\n", i, j)
	}
}
