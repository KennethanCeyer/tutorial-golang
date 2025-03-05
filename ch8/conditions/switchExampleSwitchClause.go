package main

import "fmt"

func getKoreanWord(engWord string) (korWord string) {
	var word string
	switch engWord {
	case "apple":
		word = "사과"
	case "banana":
		word = "바나나"
	case "grape":
		word = "포도"
	default:
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
