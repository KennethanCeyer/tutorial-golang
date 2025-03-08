package main

import "fmt"

type Juice struct {
	Name     string
	Price    int
	Category string
	Taste    Taste
	State    State
}

func NewJuice(
	name string,
	price int,
	category string,
	taste Taste,
	state State) *Juice {
	juice := new(Juice)
	juice.Name = name
	juice.Price = price
	juice.Category = category
	juice.Taste = taste
	juice.State = state
	return juice
}

// 쥬스를 만듦
func (j *Juice) Make() error {
	if j.State == MakeDone {
		return fmt.Errorf("%s 쥬스는 이미 만들어져 있습니다.", j.Name)
	} else if j.State == Done {
		return fmt.Errorf(
			"%s 쥬스는 이미 만들어져 고객에게 서빙되었습니다.", j.Name)
	}
	j.State = MakeDone
	return nil
}

// 쥬스를 포장함
func (j *Juice) Package() error {
	if j.State == Waiting {
		return fmt.Errorf("%s 쥬스는 아직 준비되지 않았습니다.", j.Name)
	} else if j.State == PackageDone {
		return fmt.Errorf("%s 쥬스는 이미 포장이 완료되었습니다.", j.Name)
	} else if j.State == Done {
		return fmt.Errorf(
			"%s 쥬스는 이미 고객에게 서빙되었습니다.", j.Name)
	}
	j.State = PackageDone
	return nil
}

// 쥬스를 서빙함
func (j *Juice) Pick() error {
	if j.State == Waiting {
		return fmt.Errorf("%s 쥬스는 아직 준비되지 않았습니다.", j.Name)
	} else if j.State == MakeDone {
		return fmt.Errorf("%s 쥬스는 아직 만들어지지 않았습니다.", j.Name)
	} else if j.State == Done {
		return fmt.Errorf(
			"%s 쥬스는 이미 고객에게 서빙되었습니다.", j.Name)
	}
	j.State = Done
	return nil
}

func (j *Juice) GetName() string {
	return j.Name
}

func (j *Juice) GetPrice() int {
	return j.Price
}

func (j *Juice) GetCategory() string {
	return j.Category
}

func (j *Juice) GetTaste() Taste {
	return j.Taste
}

func (j *Juice) GetState() State {
	return j.State
}
