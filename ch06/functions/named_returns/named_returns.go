package main

import (
	"fmt"
	"time"
)

func getAgeDetails(bornYear int) (age int, ageRange int) {
	// 현재 년도를 가져옴
	currentYear := time.Now().Year()

	// 현재 연도에서 태어난 년도를 빼서 나이를 구한다.
	// 한국식 나이는 출생년도에 1살부터 시작하기 때문에, 1을 더한다.
	ageValue := currentYear - bornYear + 1

	// 연령대는 2x세, 3x세 등에서 가장 앞의 숫자대에
	// 그 자릿수에 맞춰 0을 붙여주는 값으로
	// 예를들어 31세는 30대, 45세는 40대가 된다.
	// 따라서 나이를 10으로 나눈 몫에 10을 곱하여 연령대를 구한다.
	ageRangeValue := int(ageValue/10) * 10

	return ageValue, ageRangeValue
}

func main() {
	// 사용자로 부터 입력 받은 값을 저장할 변수를 선언
	var bornYear int

	// 사용자에게 정수형 값을 입력 받고,
	// 이전에 선언한 bornYear 변수에 대입한다.
	fmt.Print("태어난 년도를 입력해주세요: ")
	fmt.Scanf("%d", &bornYear)

	// getAgeDetails에 태어난 년도 정수형 변수를 인자로 전달하여 호출하고,
	// 반환된 나이와 연령대 값을 각각 age, ageRange 변수에 대입한다.
	age, ageRange := getAgeDetails(bornYear)

	// 이전에 사용자가 입력한 값과 동일한 줄에 출력되지 않도록
	// 줄넘김 한다.
	fmt.Print("\n")

	// 나이와 연령대 값을 출력한다.
	fmt.Printf("현재 나이는: %d, 연령대는 %d대 이시군요!\n",
		age, ageRange)
}
