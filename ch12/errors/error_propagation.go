package main

import (
	"fmt"
)

type ValueError struct {
    message string
}

func (e ValueError) Error() string {
    return e.message
}

func validateNumber(number int) error {
    if number < 0 {
        return ValueError{"숫자는 음수일 수 없습니다."}
    }
    return nil
}

func processNumber(number int) error {
    err := validateNumber(number)
    if err != nil {
        return fmt.Errorf("processNumber에서 오류 발생: %w", err)
    }
    return nil
}

func main() {
    numbers := []int{5, -1, 7}
    for _, number := range numbers {
        err := processNumber(number)
        if err != nil {
            fmt.Println(err)
        }
    }
}
