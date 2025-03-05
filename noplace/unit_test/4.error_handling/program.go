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

func GetTimezone(location string) (*Timezone, error) {
	if timezone[location] == nil {
		return nil, fmt.Errorf("'%s'는 지원되는 지역이 아닙니다.",
			location)
	}
	return timezone[location], nil
}

func main() {
	location := "Europe/London"
	tz, err := GetTimezone(location)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v timezone UTC offset = %d\n", location, tz.UTC)
	}
}
