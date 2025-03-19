package main

import "fmt"

func getTranslatedWord(
	engWord string,
	language string) (translatedWord string) {

	if language == "en" {
		return engWord
	} else if engWord == "apple" && language == "ko" {
		return "사과"
	} else if engWord == "apple" && language == "jp" {
		return "林檎"
	} else if engWord == "banana" && language == "ko" {
		return "바나나"
	} else if engWord == "banana" && language == "jp" {
		return "バナナ"
	} else if engWord == "grape" && language == "ko" {
		return "포도"
	} else if engWord == "grape" && language == "jp" {
		return "葡萄"
	} else {
		return "no matched word!"
	}
}

func main() {
	var word, language string

	fmt.Print("찾고 싶은 단어를 입력해주세요 (apple, banana, grape): ")
	fmt.Scanf("%s", &word)

	fmt.Print("변환할 언어를 입력해주세요 (지원되는 언어: ko, jp, en): ")
	fmt.Scanf("%s", &language)

	translatedWord := getTranslatedWord(word, language)
	fmt.Print("\n")
	fmt.Printf("찾은 단어: %v\n", translatedWord)
}
