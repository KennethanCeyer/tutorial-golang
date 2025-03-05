package main


import (
    "log"
    "os"
)


func main() {
    file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatal("파일을 열 수 없습니다: ", err)
    }
    log.SetOutput(file)
    log.Println("파일에 로그를 기록합니다.")
}

