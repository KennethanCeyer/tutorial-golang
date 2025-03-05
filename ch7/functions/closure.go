package main

import "fmt"

// 바깥 함수
func GetIncreaser() func() {
    value := 100

    // 안쪽 함수
    return func() {
        // 바깥 함수의 value 변수 접근
        value += 50
        fmt.Printf("value: %d\n", value)
    }
}

func main() {
    increaser := GetIncreaser()
    for i := 0; i < 5; i++ {
        increaser()
    }
}
