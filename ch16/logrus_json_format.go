package main


import (
    log "github.com/sirupsen/logrus"
)


func main() {
    log.SetFormatter(&log.JSONFormatter{})
    log.Info("JSON 형식의 로그 메시지입니다.")
}

