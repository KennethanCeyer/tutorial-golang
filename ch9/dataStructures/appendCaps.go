package main

import "fmt"

func main() {
	// nil 슬라이스 선언
	var sliceValue []int

	fmt.Println("================================================")
	fmt.Printf("요소 추가 전 sliceValue의 len: %d, cap: %d\n",
		len(sliceValue), cap(sliceValue))
	fmt.Println("================================================")

	var itemNumber int
	fmt.Print("배열에 요소 몇개를 더 할까요?: ")
	fmt.Scanf("%d", &itemNumber)

	for i := 0; i < itemNumber; i++ {
		// append 함수를 이용해서 슬라이스에 요소 추가
		sliceValue = append(sliceValue, i)

		fmt.Printf("%d번 요소 추가, len: %d, cap: %d\n",
			i, len(sliceValue), cap(sliceValue))
	}

	fmt.Println("================================================")
	fmt.Printf("요소 추가 후 sliceValue의 len: %d, cap: %d\n",
		len(sliceValue), cap(sliceValue))
	fmt.Println("================================================")
}
