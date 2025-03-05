package main


import (
    "context"
    "fmt"
)


func main() {
    ctx := context.WithValue(context.Background(), "userID", 12345)


    userID := ctx.Value("userID").(int)
    fmt.Println("사용자 ID:", userID)
}

