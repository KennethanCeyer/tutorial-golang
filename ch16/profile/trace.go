package main

import (
	"fmt"
	"log"
	_ "net/http/pprof"
	"os"
	"runtime/trace"
)

func main() {
    // 트레이스 파일 생성
    f, err := os.Create("trace.out")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    // 트레이스 시작
    if err := trace.Start(f); err != nil {
        log.Fatal(err)
    }
    defer trace.Stop()

    // 시뮬레이션 작업
    for i := 0; i < 10; i++ {
        fmt.Println(i)
    }
}
