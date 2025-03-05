package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) Walk() {
	fmt.Println("걷고 있습니다.")
}

func (p Person) Talk() {
	fmt.Printf("안녕하세요 제 이름은 %s 입니다.\n", p.Name)
}

type Student struct {
	Person
}

func (s Student) Study() {
	fmt.Println("공부하고 있습니다.")
}

type Engineer struct {
	Person
}

func (e Engineer) Develop() {
	fmt.Println("개발하고 있습니다.")
}

type Reporter struct {
	Person
}

func (r Reporter) Talk() {
	fmt.Printf("안녕하세요 저는 보도국 기자 %s 입니다.\n", r.Name)
}

func (r Reporter) Publish(news string) {
	fmt.Printf("%s라는 최신 보도가 있습니다.", news)
}

func main() {
	person := Person{Name: "홍길동"}
	student := Student{Person: Person{Name: "김철수"}}
	engineer := Engineer{Person: Person{Name: "리누스 토발즈"}}
	reporter := Reporter{Person: Person{Name: "박민수"}}

	person.Talk()
	student.Talk()
	engineer.Talk()
	reporter.Talk()
}
