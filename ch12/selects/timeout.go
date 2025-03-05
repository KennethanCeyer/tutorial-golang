package main

import (
	"fmt"
	"time"
)


func sendValue(value int, ch chan int) {
    time.Sleep(time.Second * 2)
    ch <- value
}


func main() {
    channel := make(chan int)
    go sendValue(1, channel)


    select {
    case value := <-channel:
        fmt.Printf("received from chan: %d\n", value)
    case <-time.After(time.Second * 1):
        fmt.Println("timeout: channel not ready")
    }


    fmt.Println("program ends")
}
