package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    // 로컬호스트의 포트 8080에 TCP 연결 시도
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("서버 연결 실패:", err)
        return
    }
    defer conn.Close()
    reader := bufio.NewReader(os.Stdin)

    for {
        // 사용자로부터 메시지 입력 받기
        fmt.Print("메시지 입력: ")
        text, _ := reader.ReadString('\n')
        // 서버로 메시지 전송
        _, err = fmt.Fprintf(conn, text)
        if err != nil {
            fmt.Println("메시지 전송 실패:", err)
            return
        }
        // 서버로부터 응답 수신
        message, err := bufio.NewReader(conn).ReadString('\n')
        if err != nil {
            fmt.Println("서버 응답 수신 실패:", err)
            return
        }
        fmt.Print("서버로부터 응답:", message)
    }
}

