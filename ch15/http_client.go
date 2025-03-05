package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    // HTTP GET 요청 보내기
    resp, err := http.Get("http://www.google.com")
    if err != nil {
        fmt.Println("HTTP 요청 실패:", err)
        return
    }
    defer resp.Body.Close()

    // 응답 본문 읽기
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("응답 읽기 실패:", err)
        return
    }
    fmt.Println("응답 상태 코드:", resp.StatusCode)
    fmt.Println("응답 본문 길이:", len(body))
}

