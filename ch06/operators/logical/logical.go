package main

import "fmt"

func main() {
	fmt.Println("AND operator")
	fmt.Println("========")
	fmt.Printf("%v\t=\ttrue && true\n", true && true)
	fmt.Printf("%v\t=\ttrue && false\n", true && false)
	fmt.Printf("%v\t=\tfalse && true\n", false && true)
	fmt.Printf("%v\t=\tfalse && true\n", false && false)
	fmt.Println("========")


	fmt.Println("OR operator")
	fmt.Println("========")
	fmt.Printf("%v\t=\ttrue || true\n", true || true)
	fmt.Printf("%v\t=\ttrue || false\n", true || false)
	fmt.Printf("%v\t=\tfalse || true\n", false || true)
	fmt.Printf("%v\t=\tfalse || true\n", false || false)
	fmt.Println("========")


	fmt.Println("NOT operator")
	fmt.Println("========")
	fmt.Printf("%v\t=\t!true\n", !true)
	fmt.Printf("%v\t=\t!false\n", !false)
}
