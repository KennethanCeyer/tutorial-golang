package main

import "fmt"

func main() {
	score := make([]int, 5)

	// 사용자로 부터 5 과목의 점수 입력
	for i, scoreItem := range score {
		fmt.Printf("%d 번째 과목 점수를 입력하세요: ", i+1)
		fmt.Scanf("%d", &scoreItem)
	}
	fmt.Println("============================")

	// 모든 슬라이스 요소를 반복하며 평균 점수 계산
	sumScore := 0
	for _, scoreItem := range score {
		sumScore += scoreItem
	}
	meanScore := float64(sumScore) / float64(len(score))
	fmt.Printf("당신의 점수 평균은 %.2f입니다.\n", meanScore)
}
