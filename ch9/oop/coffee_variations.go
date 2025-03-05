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
	americano := NewCoffee("아메리카노", 3000, "블랜딩 커피", Bitter)
	latte := NewCoffee("카페라떼", 3500, "블랜딩 커피", Sweet)
	cateMocha := NewCoffee("카페모카", 4000, "디저트 커피", Sweet)
	dripCoffe := NewCoffee("드립커피", 7000, "원두 커피", FruitFlavor)
	dutchCoffee := NewCoffee("더치커피", 5000, "더치 커피", Bitter)

	fmt.Printf("americano: %v\n", americano)
	fmt.Printf("latte: %v\n", latte)
	fmt.Printf("cateMocha: %v\n", cateMocha)
	fmt.Printf("dripCoffe: %v\n", dripCoffe)
	fmt.Printf("dutchCoffee: %v\n", dutchCoffee)
}
