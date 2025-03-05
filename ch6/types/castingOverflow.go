package main

import (
	"fmt"
	"math"
)

func main() {
	// int64alue 변수는 int64에서 표현 가능한 최대 값을 할당
	var int64Value int64 = math.MaxInt64

	// int32Value 변수에 int32로 캐스팅된 int64Value 값을 할당
	// int32 자료형의 표현 범위보다 큰 값을 캐스팅하였기에
	// 오버플로우(Overflow) 문제가 발생함
	var int32Value int32 = int32(int64Value)

	fmt.Printf("int32 value: %d\n", int32Value)
	fmt.Printf("int64 value: %d\n", int64Value)
}
