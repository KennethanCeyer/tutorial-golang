package main

import (
    "bufio"
    "fmt"
    "net"
)

func main() {
    // 포트 8080에서 TCP 연결을 수신 대기
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("서버 시작 실패:", err)
        return
    }
    defer listener.Close()
    fmt.Println("TCP 서버가 포트 8080에서 실행 중입니다.")

    for {
        // 클라이언트의 연결 요청 수락
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("연결 수락 실패:", err)
            continue
        }
        // 각 연결을 고루틴으로 처리하여 동시에 여러 클라이언트를 처리할 수 있게 함
        go handleConnection(conn)
    }
}

// 클라이언트 연결 처리 함수
func handleConnection(conn net.Conn) {
    defer conn.Close()
    reader := bufio.NewReader(conn)

    for {
        // 클라이언트로부터 메시지 읽기
        message, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("클라이언트 연결 종료:", err)
            return
        }
        fmt.Print("수신한 메시지:", message)
        // 수신한 메시지를 클라이언트에게 다시 전송 (에코)
        _, err = conn.Write([]byte("서버 응답: " + message))
        if err != nil {
            fmt.Println("메시지 전송 실패:", err)
            return
        }
    }
}

