package main

import "fmt"

type Rectangle struct {
	x, y int
}

func NewRectangle() *Rectangle {
	return new(Rectangle)
}

func (r Rectangle) Draw() {
	fmt.Println("사각형을 그립니다!")
}

func main() {
	rectangle := NewRectangle()
	rectangle.Draw()
}
