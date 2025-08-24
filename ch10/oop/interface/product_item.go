package main

// Coffee, Juice, Tea 객체의 인터페이스
// Make(), Package(), Pick() 메소드 정의
type ProductItem interface {
	Make() error
	Package() error
	Pick() error

	// 멤버 변숫값을 가져오는 게터 함수 정의
	Name() string
	Price() int
	Category() string
	Taste() Taste
	State() State
}
