package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d번째 반복문\n", i)
	}
}
