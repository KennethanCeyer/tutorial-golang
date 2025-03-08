package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var nameSlice []string
	var name string

	for i := 0; name != "exit"; i++ {
		fmt.Printf("%d번째 사람의 이름을 입력해주세요: ", i+1)
		fmt.Scanf("%s", &name)
		nameSlice = append(nameSlice, name)
	}

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(len(nameSlice))

	fmt.Print("\n")
	fmt.Printf("이번 점심값은 %s님이 쏩니다!\n", nameSlice[randomNumber])
}
