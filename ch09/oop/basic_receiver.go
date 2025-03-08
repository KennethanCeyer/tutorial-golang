package main

import "fmt"

type Building struct {
	Status string
}

func NewBuilding() *Building {
	return new(Building)
}

// Building 객체의 메소드 Open 정의
func (b *Building) Open() {
	b.Status = "OPEN"
}

func main() {
	building := NewBuilding()
	fmt.Printf("building status: %s\n", building.Status)

	building.Open()
	fmt.Printf("building status: %s\n", building.Status)
}
