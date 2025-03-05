package main

// Coffee, Juice, Tea 객체의 인터페이스
// Make(), Package(), Pick() 메소드 정의
type ProductItem interface {
	Make() error
	Package() error
	Pick() error

	// 멤버 필드를 가져오는 Getter 함수 정의
	GetName() string
	GetPrice() int
	GetCategory() string
	GetTaste() Taste
	GetState() State
}
