package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		n := i + 1
		fmt.Printf("%d번째 반복문: i=%d\n", n, i)
	}
}
