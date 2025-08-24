package main

import "fmt"

type Juice struct {
	name     string
	price    int
	category string
	taste    Taste
	state    State
}

func NewJuice(
	name string,
	price int,
	category string,
	taste Taste,
	state State,
) *Juice {
	return &Juice{name: name, price: price, category: category, taste: taste, state: state}
}

// 주스를 만듦
func (j *Juice) Make() error {
	if j.state == MakeDone {
		return fmt.Errorf("%s 주스는 이미 만들어져 있습니다.", j.name)
	} else if j.state == Done {
		return fmt.Errorf(
			"%s 주스는 이미 만들어져 고객에게 서빙되었습니다.", j.name)
	}
	j.state = MakeDone
	return nil
}

// 주스를 포장함
func (j *Juice) Package() error {
	if j.state == Waiting {
		return fmt.Errorf("%s 주스는 아직 준비되지 않았습니다.", j.name)
	} else if j.state == PackageDone {
		return fmt.Errorf("%s 주스는 이미 포장이 완료되었습니다.", j.name)
	} else if j.state == Done {
		return fmt.Errorf(
			"%s 주스는 이미 고객에게 서빙되었습니다.", j.name)
	}
	j.state = PackageDone
	return nil
}

// 주스를 서빙함
func (j *Juice) Pick() error {
	if j.state == Waiting {
		return fmt.Errorf("%s 주스는 아직 준비되지 않았습니다.", j.name)
	} else if j.state == MakeDone {
		return fmt.Errorf("%s 주스는 아직 포장되지 않았습니다.", j.name)
	} else if j.state == Done {
		return fmt.Errorf(
			"%s 주스는 이미 고객에게 서빙되었습니다.", j.name)
	}
	j.state = Done
	return nil
}

func (j *Juice) Name() string {
	return j.name
}

func (j *Juice) Price() int {
	return j.price
}

func (j *Juice) Category() string {
	return j.category
}

func (j *Juice) Taste() Taste {
	return j.taste
}

func (j *Juice) State() State {
	return j.state
}
