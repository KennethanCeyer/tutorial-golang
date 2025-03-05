package main

import (
    "log"

    "gopkg.in/natefinch/lumberjack.v2"
)

func main() {
    log.SetOutput(&lumberjack.Logger{
        Filename:   "app.log",
        MaxSize:    5, // 메가바이트 단위
        MaxBackups: 3,
        MaxAge:     28,   // 일 단위
        Compress:   true, // 압축 여부
    })

    for i := 0; i < 1000000; i++ {
        log.Printf("로그 메시지 %d", i)
    }
}

