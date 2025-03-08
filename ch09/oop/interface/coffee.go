package main

import "fmt"

type Coffee struct {
	Name     string
	Price    int
	Category string
	Taste    Taste
	State    State
}

func NewCoffee(
	name string,
	price int,
	category string,
	taste Taste,
	state State) *Coffee {
	coffee := new(Coffee)
	coffee.Name = name
	coffee.Price = price
	coffee.Category = category
	coffee.Taste = taste

	// 추가됨
	coffee.State = state
	return coffee
}

// 커피를 만듦
func (c *Coffee) Make() error {
	if c.State == MakeDone {
		return fmt.Errorf("%s 커피는 이미 만들어져 있습니다.", c.Name)
	} else if c.State == Done {
		return fmt.Errorf(
			"%s 커피는 이미 만들어져 고객에게 서빙되었습니다.", c.Name)
	}
	c.State = MakeDone
	return nil
}

// 커피를 포장함
func (c *Coffee) Package() error {
	if c.State == Waiting {
		return fmt.Errorf("%s 커피는 아직 준비되지 않았습니다.", c.Name)
	} else if c.State == PackageDone {
		return fmt.Errorf("%s 커피는 이미 포장이 완료되었습니다.", c.Name)
	} else if c.State == Done {
		return fmt.Errorf(
			"%s 커피는 이미 고객에게 서빙되었습니다.", c.Name)
	}
	c.State = PackageDone
	return nil
}

// 커피를 서빙함
func (c *Coffee) Pick() error {
	if c.State == Waiting {
		return fmt.Errorf("%s 커피는 아직 준비되지 않았습니다.", c.Name)
	} else if c.State == MakeDone {
		return fmt.Errorf("%s 커피는 아직 만들어지지 않았습니다.", c.Name)
	} else if c.State == Done {
		return fmt.Errorf(
			"%s 커피는 이미 고객에게 서빙되었습니다.", c.Name)
	}
	c.State = Done
	return nil
}

func (c *Coffee) GetName() string {
	return c.Name
}

func (c *Coffee) GetPrice() int {
	return c.Price
}

func (c *Coffee) GetCategory() string {
	return c.Category
}

func (c *Coffee) GetTaste() Taste {
	return c.Taste
}

func (c *Coffee) GetState() State {
	return c.State
}
