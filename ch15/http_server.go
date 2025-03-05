package main

import (
    "fmt"
    "net/http"
)

// 루트 경로 요청 처리 함수
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "안녕하세요, Go로 만든 웹 서버입니다!")
}

func main() {
    // 요청 경로와 핸들러 함수 매핑
    http.HandleFunc("/", helloHandler)
    fmt.Println("HTTP 서버가 포트 8080에서 실행 중입니다.")
    // 서버 실행
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("서버 실행 실패:", err)
    }
}

