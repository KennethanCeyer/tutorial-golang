package main

import (
	"os"
)


func openFile(filePath string) {
    _, err := os.OpenFile(filePath, os.O_RDONLY, os.FileMode(0644))
    if err != nil {
        panic(err)  // 파일을 열 수 없으면 패닉 상태로 전환한다.
    }
}


func main() {
    filePath := "noFile"
    openFile(filePath)  // 1존재하지 않는 파일을 열려고 시도하여 패닉을 발생시킨다.
}
