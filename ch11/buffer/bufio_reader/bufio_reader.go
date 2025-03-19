package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("문장을 입력하세요: ")
	text, _ := reader.ReadString('\n') // 개행 문자 포함 입력 읽기
	text = strings.TrimSpace(text)     // 불필요한 공백 제거
	fmt.Printf("입력된 문장: %s\n", text)
}
