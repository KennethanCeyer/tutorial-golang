package main


import (
    "log"
)


func main() {
    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
    log.Println("로그 설정이 변경되었습니다.")
}

