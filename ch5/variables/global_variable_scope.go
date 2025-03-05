package main

import "fmt"

var (
	message = "hello"
)

func display() {
	fmt.Println(message)
}

func main() {
	message += " world"
	display()
}
