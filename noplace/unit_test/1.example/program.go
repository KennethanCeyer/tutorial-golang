package main

import "fmt"

type Timezone struct {
	Name string
	UTC  int
}

var (
	timezone = map[string]*Timezone{
		"Asia/Seoul":       &Timezone{Name: "KST", UTC: 9},
		"Asia/Tokyo":       &Timezone{Name: "JST", UTC: 9},
		"America/New_York": &Timezone{Name: "EST", UTC: -5},
	}
)

func main() {
	location := "Asia/Seoul"
	tz := timezone[location]
	fmt.Printf("%v timezone UTC offset = %d\n", location, tz.UTC)
}
