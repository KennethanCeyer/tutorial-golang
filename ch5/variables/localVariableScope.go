package main

import "fmt"

func display() {
	fmt.Println(message)
}

func main() {
	var message string = "hello"
	message += " world"
	display()
}
