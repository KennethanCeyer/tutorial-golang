package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)


func populateChan(ch chan int, wg *sync.WaitGroup) {
    for i := 0; i < 10; i++ {
        ch <- rand.Int()
        time.Sleep(50 * time.Millisecond)
    }
    wg.Done()
}


func display(ch1 chan int, ch2 chan int) {
    for {
        select {
        case value := <-ch1:
            fmt.Printf("ch1: %d\n", value)
        case value := <-ch2:
            fmt.Printf("ch2: %d\n", value)
        default:
            // 아무 채널도 준비되지 않았을 때의 동작
        }
    }
}


func main() {
    rand.Seed(time.Now().UnixNano())
    wg := &sync.WaitGroup{}


    channel1 := make(chan int)
    channel2 := make(chan int)
    wg.Add(2)
    go display(channel1, channel2)
    go populateChan(channel1, wg)
    go populateChan(channel2, wg)
    
    // 채널에 값을 완전히 넣을 때까지 기다리고, 종료
    wg.Wait()
}
