package main

import "fmt"

func getScoreRank(score int) (rank string) {
	if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else {
		return "F"
	}
}

func main() {
	var score int

	fmt.Print("점수를 입력해주세요: ")
	fmt.Scanf("%d", &score)

	rank := getScoreRank(score)
	fmt.Printf("당신의 점수 등급은 '%v' 입니다.\n", rank)
}
