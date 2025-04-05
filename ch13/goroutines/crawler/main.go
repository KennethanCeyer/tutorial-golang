package main

import (
	"fmt"
	"strings"
	"time"
)

const divideBar = "================================================"

func main() {
	var confirm string
	useGoroutine := false
	fmt.Print("고루틴을 사용하시겠습니까? (y/N): ")
	fmt.Scanf("%s", &confirm)

	if strings.ToUpper(confirm) == "Y" {
		useGoroutine = true
	}

	index, err := Indexing()
	if err != nil {
		panic(err)
	}

	fmt.Println(divideBar)
	fmt.Printf("총 %d건의 뉴스 기사를 찾았습니다.\n", len(index.Link))
	fmt.Println("뉴스 내용을 불러옵니다.")

	start := time.Now()
	err = Crawler(index, useGoroutine)
	if err != nil {
		panic(err)
	}
	elapsed := time.Since(start)
	fmt.Println("모든 뉴스 내용을 불러왔습니다.")
	fmt.Println(divideBar)
	for _, link := range index.Link {
		fmt.Printf("TITLE  : %v\n", link.Title)
		fmt.Printf("LINK   : %v\n", link.URI)
		fmt.Printf("CONTENT: %d 글자의 내용\n", len(link.Content))
		fmt.Println(divideBar)
	}
	fmt.Printf("소요시간: %v\n", elapsed)
}
