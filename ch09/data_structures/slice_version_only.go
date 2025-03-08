package main

import (
	"fmt"
	"math/rand"
)
 
func main() {
	var nameSlice []string
	var name string

	for i := 0; name != "exit"; i++ {
		fmt.Printf("%d번째 사람의 이름을 입력해주세요: ", i + 1)
		fmt.Scanf("%s", &name)
		nameSlice = append(nameSlice, name)
	}

	randomNumber := rand.Intn(len(nameSlice))

	fmt.Print("\n")
	fmt.Printf("이번 점심값은 %s님이 냅니다!\n", nameSlice[randomNumber])
 }
