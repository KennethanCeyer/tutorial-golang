package main

import "fmt"

func main() {
	increase := 1
	increase++ // increase = increase + 1과 같음
	fmt.Println(increase)

	decrease := 1
	decrease-- // decrease = decrease - 1과 같음
	fmt.Println(decrease)
}
