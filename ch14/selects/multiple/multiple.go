package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func populateChan(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- rand.Intn(1000)
		time.Sleep(50 * time.Millisecond)
	}
}

func display(ch1, ch2 <-chan int) {
	ch1Open, ch2Open := true, true

	// 두 채널 중 하나라도 열려 있으면 계속 반복
	for ch1Open || ch2Open {
		select {
		case value, ok := <-ch1:
			if !ok {
				ch1Open = false // ch1 닫힘 감지
			} else {
				fmt.Printf("ch1: %d\n", value)
			}
		case value, ok := <-ch2:
			if !ok {
				ch2Open = false // ch2 닫힘 감지
			} else {
				fmt.Printf("ch2: %d\n", value)
			}
		}
	}
	fmt.Println("모든 채널이 닫혀 display 함수를 종료합니다.")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup

	channel1 := make(chan int)
	channel2 := make(chan int)

	go display(channel1, channel2)

	wg.Add(2)
	go populateChan(channel1, &wg)
	go populateChan(channel2, &wg)

	// populateChan 고루틴들이 모두 끝날 때까지 대기
	wg.Wait()

	// 데이터 생성이 완료되었으므로 채널을 닫아 display 고루틴에 알림
	close(channel1)
	close(channel2)

	time.Sleep(300 * time.Millisecond) // display() 함수의 출력을 위해 잠시 대기
	fmt.Println("main 함수 종료.")
}
