package main

import "fmt"

// 제네릭 리스트 타입 정의
type List[T any] struct {
    items []T
}


// 리스트에 항목 추가
func (l *List[T]) Add(item T) {
    l.items = append(l.items, item)
}


// 리스트 항목 출력
func (l *List[T]) Print() {
    for _, item := range l.items {
        fmt.Println(item)
    }
}


func main() {
    // 정수 타입 리스트
    intList := List[int]{}
    intList.Add(1)
    intList.Add(2)
    intList.Print()  // 출력: 1 2


    // 문자열 타입 리스트
    strList := List[string]{}
    strList.Add("hello")
    strList.Add("world")
    strList.Print()  // 출력: hello world
}
