package main

import "fmt"

type ValueError struct {
	message string
}

func (e ValueError) Error() string {
	return e.message
}

func checkBoolString(boolString string) error {
	if boolString == "true" || boolString == "false" {
		return nil
	}

	return ValueError{
		fmt.Sprintf("'%s'는 논리형 문자열이 아닙니다.", boolString)}
}

func main() {
	values := []string{"true", "false", "hello"}
	for _, value := range values {
		err := checkBoolString(value)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
