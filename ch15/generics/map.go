package main

import "fmt"

// 제네릭 맵 타입 정의
type Map[K comparable, V any] struct {
    data map[K]V
}

// 맵 생성
func NewMap[K comparable, V any]() *Map[K, V] {
    return &Map[K, V]{data: make(map[K]V)}
}

// 키-값 쌍 추가
func (m *Map[K, V]) Set(key K, value V) {
    m.data[key] = value
}

// 값 조회
func (m *Map[K, V]) Get(key K) (V, bool) {
    value, exists := m.data[key]
    return value, exists
}

func main() {
    // 문자열 키와 정수 값을 사용하는 맵
    stringIntMap := NewMap[string, int]()
    stringIntMap.Set("a", 1)
    stringIntMap.Set("b", 2)
    fmt.Println(stringIntMap.Get("a"))  // 출력: 1 true
    fmt.Println(stringIntMap.Get("b"))  // 출력: 2 true

    // 정수 키와 문자열 값을 사용하는 맵
    intStringMap := NewMap[int, string]()
    intStringMap.Set(1, "hello")
    intStringMap.Set(2, "world")
    fmt.Println(intStringMap.Get(1))  // 출력: hello true
    fmt.Println(intStringMap.Get(2))  // 출력: world true
}
