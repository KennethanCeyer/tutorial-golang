package main

import (
    "fmt"
    "net"
)

func main() {
    // TCP 프로토콜로 google.com의 80번 포트에 연결 시도
    conn, err := net.Dial("tcp", "google.com:80")
    if err != nil {
        fmt.Println("연결 실패:", err)
        return
    }
    defer conn.Close() // 프로그램 종료 시 연결 닫기

    fmt.Println("연결 성공:", conn.RemoteAddr())
}

