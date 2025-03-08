package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nameSlice := make([]string, 5)

	for i := 0; i < 5; i++ {
		fmt.Printf("%d번째 사람의 이름을 입력해주세요: ", i+1)
		fmt.Scanf("%s", &nameSlice[i])
	}

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(len(nameSlice))

	fmt.Print("\n")
	fmt.Printf("이번 점심값은 %s님이 쏩니다!\n", nameSlice[randomNumber])
}
