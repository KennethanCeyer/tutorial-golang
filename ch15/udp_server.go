package main

import (
    "fmt"
    "net"
)

func main() {
    // UDP 주소 생성
    addr, err := net.ResolveUDPAddr("udp", ":8080")
    if err != nil {
        fmt.Println("주소 해석 실패:", err)
        return
    }
    // UDP 연결 수신 대기
    conn, err := net.ListenUDP("udp", addr)
    if err != nil {
        fmt.Println("UDP 서버 시작 실패:", err)
        return
    }
    defer conn.Close()
    fmt.Println("UDP 서버가 포트 8080에서 실행 중입니다.")

    buffer := make([]byte, 1024)
    for {
        // 클라이언트로부터 데이터 수신
        n, clientAddr, err := conn.ReadFromUDP(buffer)
        if err != nil {
            fmt.Println("데이터 수신 실패:", err)
            continue
        }
        message := string(buffer[:n])
        fmt.Printf("수신한 메시지: %s", message)
        // 클라이언트에게 응답 전송
        _, err = conn.WriteToUDP([]byte("서버 응답: "+message), clientAddr)
        if err != nil {
            fmt.Println("응답 전송 실패:", err)
        }
    }
}

