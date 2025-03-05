package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	// 표시할 로그 범위 지정
	log.SetLevel(log.DebugLevel)

	// 로그 타입 별 메시지 출력
	log.WithFields(log.Fields{
		"language": "go",
		"mode":     "test",
	}).Info("안내 메시지입니다.")
	log.Debug("디버그용 출력")
	log.Error("에러 발생")
	log.Fatal("치명적인 문제 발생")
}
