package main

import "fmt"

func main() {
	unbufferedChannel := make(chan int)
	unbufferedChannel <- 1
	fmt.Println(<-unbufferedChannel)
}
