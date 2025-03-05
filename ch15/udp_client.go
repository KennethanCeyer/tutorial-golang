package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    // 서버 주소 해석
    serverAddr, err := net.ResolveUDPAddr("udp", "localhost:8080")
    if err != nil {
        fmt.Println("주소 해석 실패:", err)
        return
    }
    // UDP 연결 생성
    conn, err := net.DialUDP("udp", nil, serverAddr)
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
        _, err = conn.Write([]byte(text))
        if err != nil {
            fmt.Println("데이터 전송 실패:", err)
            continue
        }
        // 서버로부터 응답 수신
        buffer := make([]byte, 1024)
        n, _, err := conn.ReadFromUDP(buffer)
        if err != nil {
            fmt.Println("응답 수신 실패:", err)
            continue
        }
        fmt.Print("서버로부터 응답:", string(buffer[:n]))
    }
}

