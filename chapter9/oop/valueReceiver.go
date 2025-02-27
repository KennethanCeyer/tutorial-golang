package main

import "fmt"

type Rectangle struct {
	x, y int
}

func NewRectangle(x int, y int) *Rectangle {
	rectangle := new(Rectangle)
	rectangle.x = x
	rectangle.y = y
	return rectangle
}

func (r Rectangle) Draw(x int, y int) {
	fmt.Printf("변경 전 x, y: %d, %d\n", r.x, r.y)
	r.x = x
	r.y = y
	fmt.Printf("변경 후 x, y: %d, %d\n", r.x, r.y)
}

func main() {
	rectangle := NewRectangle(5, 5)
	rectangle.Draw(10, 20)
	fmt.Printf("원본 Rectangle x, y: %d, %d\n",
		rectangle.x, rectangle.y)
}
