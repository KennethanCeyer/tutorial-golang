package main

import "fmt"

func namedFunction(a int, b int) (sum int) {
	return a + b
}

func main() {
	namedAddResult := namedFunction(2, 3)
	anonymousAddResult := func(a int, b int) (sum int) {
		return a + b
	}(2, 3)

	fmt.Printf("namedAdd result: %v\n", namedAddResult)
	fmt.Printf("anonymousAdd result: %v\n", anonymousAddResult)
}
