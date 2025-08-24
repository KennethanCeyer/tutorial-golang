package main

import "fmt"

type Tea struct {
	name     string
	price    int
	category string
	taste    Taste
	state    State
}

func NewTea(
	name string,
	price int,
	category string,
	taste Taste,
	state State,
) *Tea {
	return &Tea{name: name, price: price, category: category, taste: taste, state: state}
}

func (t *Tea) Make() error {
	if t.state == MakeDone {
		return fmt.Errorf("%s 차는 이미 만들어져 있습니다.", t.name)
	} else if t.state == Done {
		return fmt.Errorf(
			"%s 차는 이미 만들어져 고객에게 서빙되었습니다.", t.name)
	}
	t.state = MakeDone
	return nil
}

// 차를 포장함
func (t *Tea) Package() error {
	if t.state == Waiting {
		return fmt.Errorf("%s 차는 아직 준비되지 않았습니다.", t.name)
	} else if t.state == PackageDone {
		return fmt.Errorf("%s 차는 이미 포장이 완료되었습니다.", t.name)
	} else if t.state == Done {
		return fmt.Errorf(
			"%s 차는 이미 고객에게 서빙되었습니다.", t.name)
	}
	t.state = PackageDone
	return nil
}

// 차를 서빙함
func (t *Tea) Pick() error {
	if t.state == Waiting {
		return fmt.Errorf("%s 차는 아직 준비되지 않았습니다.", t.name)
	} else if t.state == MakeDone {
		return fmt.Errorf("%s 차는 아직 포장되지 않았습니다.", t.name)
	} else if t.state == Done {
		return fmt.Errorf(
			"%s 차는 이미 고객에게 서빙되었습니다.", t.name)
	}
	t.state = Done
	return nil
}

func (t *Tea) Name() string {
	return t.name
}

func (t *Tea) Price() int {
	return t.price
}

func (t *Tea) Category() string {
	return t.category
}

func (t *Tea) Taste() Taste {
	return t.taste
}

func (t *Tea) State() State {
	return t.state
}
