package main

import (
	"book_api/controllers"
	"book_api/database"
	"book_api/repositories"
	"book_api/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
)

func init() {
    file, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatal("로그 파일을 열 수 없습니다:", err)
    }
    log.SetOutput(file)
    log.Println("서버 시작")
}

func main() {
    // 데이터베이스 초기화
    db, err := database.InitDB(sqlite.Open("books.db"))
    if err != nil {
        log.Fatalf("데이터베이스 초기화 실패: %v", err)
    }

    // Gin 엔진 생성
    r := gin.New()

    // 로깅 및 에러 복구 미들웨어 추가
    r.Use(gin.LoggerWithWriter(log.Writer()))
    r.Use(gin.Recovery())

    // 의존성 주입
    repo := &repositories.GormBookRepository{DB: db}
    bookController := &controllers.BookController{Repository: repo}

    // 라우트 설정
    routes.SetupRoutes(r, bookController)

    // 서버 실행
    r.Run(":8080")
}
