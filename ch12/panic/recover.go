package main

import (
	"fmt"
	"os"
)


func appendText() {
    f, _ := os.OpenFile(
        "/tmp/appendFile", os.O_WRONLY|os.O_CREATE|os.O_APPEND,
        os.FileMode(0644))


    defer func() {
        fmt.Println("defer 발생!")
        f.Close()
        if r := recover(); r != nil {
            fmt.Printf("패닉 상태 복구! recover: %v\n", r)
        }
    }()


    panic("의도적으로 패닉 발생")
}


func main() {
    appendText()
}
