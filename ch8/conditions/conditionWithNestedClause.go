package main

import "fmt"

func getTranslatedWord(
	engWord string,
	language string) (translatedWord string) {

	if language == "en" {
		return engWord
	} else if language == "ko" {
		if engWord == "apple" {
			return "사과"
		} else if engWord == "banana" {
			return "바나나"
		} else if engWord == "grape" {
			return "포도"
		} else {
			return "no matched word!"
		}
	} else if language == "jp" {
		if engWord == "apple" {
			return "林檎"
		} else if engWord == "banana" {
			return "バナナ"
		} else if engWord == "grape" {
			return "葡萄"
		} else {
			return "no matched word!"
		}
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
