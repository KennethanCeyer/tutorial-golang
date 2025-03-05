package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	// 파일 열기
	file, err := os.OpenFile("example.txt", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	
	// 파일 디스크립터 가져오기
	fd := int(file.Fd())
	fmt.Printf("File Descriptor: %d\n", fd)
	
	// 파일의 상태 확인
	info, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Println("File Name:", info.Name())
	
	// 시스템 호출을 통한 파일 상태 확인
	var stat syscall.Stat_t
	if err := syscall.Fstat(fd, &stat); err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Printf("File Size: %d bytes\n", stat.Size)
}
