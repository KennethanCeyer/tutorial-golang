package main

import "fmt"

type Tea struct {
	Name     string
	Price    int
	Category string
	Taste    Taste
	State    State
}

func NewTea(
	name string,
	price int,
	category string,
	taste Taste,
	state State) *Tea {
	tea := new(Tea)
	tea.Name = name
	tea.Price = price
	tea.Category = category
	tea.Taste = taste
	tea.State = state
	return tea
}

// 차를 만듦
func (t *Tea) Make() error {
	if t.State == MakeDone {
		return fmt.Errorf("%s 차는 이미 만들어져 있습니다.", t.Name)
	} else if t.State == Done {
		return fmt.Errorf(
			"%s 차는 이미 만들어져 고객에게 서빙되었습니다.", t.Name)
	}
	t.State = MakeDone
	return nil
}

// 차를 포장함
func (t *Tea) Package() error {
	if t.State == Waiting {
		return fmt.Errorf("%s 차는 아직 준비되지 않았습니다.", t.Name)
	} else if t.State == PackageDone {
		return fmt.Errorf("%s 차는 이미 포장이 완료되었습니다.", t.Name)
	} else if t.State == Done {
		return fmt.Errorf(
			"%s 차는 이미 고객에게 서빙되었습니다.", t.Name)
	}
	t.State = PackageDone
	return nil
}

// 차를 서빙함
func (t *Tea) Pick() error {
	if t.State == Waiting {
		return fmt.Errorf("%s 차는 아직 준비되지 않았습니다.", t.Name)
	} else if t.State == MakeDone {
		return fmt.Errorf("%s 차는 아직 만들어지지 않았습니다.", t.Name)
	} else if t.State == Done {
		return fmt.Errorf(
			"%s 차는 이미 고객에게 서빙되었습니다.", t.Name)
	}
	t.State = Done
	return nil
}

func (t *Tea) GetName() string {
	return t.Name
}

func (t *Tea) GetPrice() int {
	return t.Price
}

func (t *Tea) GetCategory() string {
	return t.Category
}

func (t *Tea) GetTaste() Taste {
	return t.Taste
}

func (t *Tea) GetState() State {
	return t.State
}
