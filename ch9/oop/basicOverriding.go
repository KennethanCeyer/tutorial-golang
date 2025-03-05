package main

import "fmt"


type Animal struct {
    Name string
}


func (a Animal) Speak() {
    fmt.Println("Animal speaks")
}


type Dog struct {
    Animal
}


func (d Dog) Speak() {
    fmt.Println("Dog barks")
}


func main() {
    animal := Animal{Name: "Generic Animal"}
    dog := Dog{Animal: Animal{Name: "Buddy"}}


    animal.Speak() // 출력: Animal speaks
    dog.Speak()    // 출력: Dog barks
}
