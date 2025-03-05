package main

import "fmt"

func main() {
loop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("i=%d, j=%d\n", i, j)
			if j == 1 {
				break loop
			}
		}
	}
}
