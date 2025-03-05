package main


import (
    "context"
    "fmt"
    "time"
)


func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()


    select {
    case <-ctx.Done():
        fmt.Println("타임아웃 발생:", ctx.Err())
    }
}

