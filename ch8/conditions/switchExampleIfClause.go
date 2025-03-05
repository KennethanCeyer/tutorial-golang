package main

import "fmt"

func getKoreanWord(engWord string) (korWord string) {
	var word string
	if engWord == "apple" {
		word = "사과"
	} else if engWord == "banana" {
		word = "바나나"
	} else if engWord == "grape" {
		word = "포도"
	} else {
		word = "no matched word!"
	}
	return word
}

func main() {
	var searchWord string

	fmt.Print("찾고 싶은 단어를 입력해주세요 (apple, banana, grape): ")
	fmt.Scanf("%s", &searchWord)

	fmt.Print("\n")
	korWord := getKoreanWord(searchWord)
	fmt.Printf("찾은 단어: %v\n", korWord)
}
