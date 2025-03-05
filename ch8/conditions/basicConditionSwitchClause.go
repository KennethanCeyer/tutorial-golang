package main

import "fmt"

func getScoreRank(score int) (rank string) {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	default:
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
