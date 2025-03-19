package main

import "fmt"

func main() {
	type Student struct {
		Name       string
		Age        int
		Department string
		Score      map[string]int
	}

	student := Student{"김철수", 20, "컴퓨터공학", map[string]int{
		"과학": 92,
		"수학": 94,
		"영어": 95,
	}}

	fmt.Printf("구조체 student: %v\n", student)
	fmt.Println("================================================")
	fmt.Printf("name: %s\n", student.Name)
	fmt.Printf("age: %d\n", student.Age)
	fmt.Printf("department: %s\n", student.Department)
	fmt.Println("================================================")
	fmt.Println("[과목 점수]")
	for name, score := range student.Score {
		fmt.Printf("%s: %d\n", name, score)
	}
	fmt.Println("================================================")
}
