package main

import (
	"fmt"
	"os"
)

func appendText() {
	f, _ := os.OpenFile(
		"/tmp/appendFile", os.O_WRONLY|os.O_CREATE|os.O_APPEND,
		os.FileMode(0644))

	// 함수가 종료될 때 반드시 File 리소스를 반환하도록 함
	defer func() {
		fmt.Println("defer 발생!")
		f.Close()
	}()

	panic("의도적으로 패닉 발생")
}

func main() {
	appendText()
}
