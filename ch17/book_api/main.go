package main

import (
	"log"
	"os"

	"book_api/database"
	"book_api/routes"

	"github.com/gin-gonic/gin"
)


func init() {
    // 로그 파일 설정
    file, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatal("로그 파일을 열 수 없습니다:", err)
    }
    log.SetOutput(file)
    log.Println("서버 시작")
}


func main() {
    database.InitDB()


    r := gin.New()


    // 로깅 미들웨어 추가
    r.Use(gin.LoggerWithWriter(log.Writer()))
    r.Use(gin.Recovery())


    // 라우트 설정
    routes.SetupRoutes(r)


    r.Run(":8080")
}
