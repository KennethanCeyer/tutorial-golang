package main

import "fmt"


type Shape interface {
    Draw()
}


type Circle struct{}


func (c Circle) Draw() {
    fmt.Println("Drawing a Circle")
}


type Square struct{}


func (s Square) Draw() {
    fmt.Println("Drawing a Square")
}


func main() {
    var shape Shape


    shape = Circle{}
    shape.Draw() // 출력: Drawing a Circle


    shape = Square{}
    shape.Draw() // 출력: Drawing a Square
}
