package main

import (
	"fmt"
	"time"
)

func timeout(sleepTime time.Duration, callback func()) {
	time.Sleep(sleepTime)
	callback()
}

func main() {
	timeout(time.Second*1, func() {
		fmt.Println("hello world!")
	})
}
