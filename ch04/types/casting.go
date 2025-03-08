package main

import (
	"fmt"
	"math"
)

func main() {
	// int64Value 변수에 int64로 캐스팅된 int32Value 값을 할당
	var int32Value int32 = math.MaxInt32

	// int64Value 변수에 int64로 캐스팅된 int32Value 값을 할당
	var int64Value int64 = int64(int32Value)

	fmt.Printf("int32 value: %d\n", int32Value)
	fmt.Printf("int64 value: %d\n", int64Value)
}
