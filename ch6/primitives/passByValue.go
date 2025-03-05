package main

import "fmt"

func updateValue(value int) {
	value = 100
}

func main() {
	myValue := 50
	fmt.Printf("myValue 변경 전: %d\n", myValue)

	updateValue(myValue)
	fmt.Printf("myValue 변경 후: %d\n", myValue)
}
