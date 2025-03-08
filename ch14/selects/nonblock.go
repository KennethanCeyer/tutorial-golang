package main

import (
	"fmt"
	"time"
)


func sendValue(value int, ch chan int) {
    time.Sleep(time.Second * 3)
    ch <- value
}


func main() {
    channel := make(chan int)
    go sendValue(1, channel)


    select {
    case value := <-channel:
        fmt.Printf("received from chan: %d\n", value)
    default:
        fmt.Println("channel is not yet prepared (unblock)")
    }


    fmt.Println("program ends")
}
