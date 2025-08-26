package main

import "fmt"

type Taste string

const (
	Sweet       Taste = "SWEET"
	Bitter      Taste = "BITTER"
	FruitFlavor Taste = "FRUIT"
	Heavy       Taste = "HEAVY"
)

type Coffee struct {
	Name     string
	Price    int
	Category string
	Taste    Taste
}

func NewCoffee(
	name string,
	price int,
	category string,
	taste Taste) *Coffee {
	coffee := new(Coffee)
	coffee.Name = name
	coffee.Price = price
	coffee.Category = category
	coffee.Taste = taste
	return coffee
}

func main() {
	coffee := NewCoffee("아메리카노", 3000, "블랜딩 커피", Bitter)
	fmt.Println("================================================")
	fmt.Printf("객체 coffee: %v\n", coffee)
	fmt.Println("================================================")
	fmt.Printf("name: %s\n", coffee.Name)
	fmt.Printf("price: %d\n", coffee.Price)
	fmt.Printf("category: %s\n", coffee.Category)
	fmt.Printf("taste: %s\n", coffee.Taste)
	fmt.Println("================================================")
}
