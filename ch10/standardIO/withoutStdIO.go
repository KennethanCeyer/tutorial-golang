package main

import (
	"fmt"
	"os"
)

func main() {
    // 현재 터미널에 직접 출력하기
    // tty 세션 파일을 직접 열어서 출력
    file, err := os.OpenFile("/dev/tty", os.O_WRONLY, 0600)
    if err != nil {
        fmt.Println("Error opening /dev/tty:", err)
        return
    }
    defer file.Close()

    // 터미널에 메시지 쓰기
    _, err = file.WriteString("Hello, this is a message directly to the terminal!\n")
    if err != nil {
        fmt.Println("Error writing to terminal:", err)
        return
    }
}
