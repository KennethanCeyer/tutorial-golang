package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.InfoLevel)

	log.Debug("이 메시지는 출력되지 않습니다.")
	log.Info("정보 메시지입니다.")
	log.Warn("경고 메시지입니다.")
	log.Error("에러 메시지입니다.")
}
