package main

import "fmt"

func getRank(score int) string {
	switch {
	case score > 90:
		return "A"
	case score > 80:
		return "B"
	case score > 70:
		return "C"
	default:
		return "F"
	}
}

func main() {
	// map 자료구조는 make(map[키 자료형]값 자료형) 형태로 할당할 수 있다
	// 키는 값을 얻기 위해 사용하는 인덱스
	// 값은 키를 인덱스로 넣었을 때 반환하는 값
	mapScore := make(map[string]int)

	for i := 0; i < 3; i++ {
		var subjectName string
		var subjectScore int

		fmt.Print("과목 이름을 입력해주세요: ")
		fmt.Scanf("%s", &subjectName)

		fmt.Printf("%s의 점수를 입력해주세요: ", subjectName)
		fmt.Scanf("%d", &subjectScore)

		mapScore[subjectName] = subjectScore
	}

	fmt.Print("\n")
	fmt.Print("============================")
	fmt.Print("\n")

	// map 자료구조를 range를 이용해서 반복시키면
	// 키와 값을 각각 name, score에 할당하여 반복한다
	for name, score := range mapScore {
		fmt.Printf("%s 과목의 등급은 %s 입니다!\n", name, getRank(score))
	}
}
