package main

import "fmt"

func main() {
    var value1 uint8 = 187    // 10111011 [187]
    var value2 uint8 = 222    // 11011110 [222]

    // 시프트 연산 값은 조금 더 보기 쉽도록 별도 변수로 저장
    var shiftValue uint8 = 19 // 00010011 [19]

    // 시프트 연산은 2를 주어 두 칸 좌 우로 이동시키도록 한다.
    var shiftLength uint8 = 2

    // AND 비트 연산
    fmt.Println("AND result")
    fmt.Println("========")
    andResult := value1 & value2 // 10011010 [154]
    fmt.Printf("%08b & %08b = %08b\n", value1, value2, andResult)

    // OR 비트 연산
    fmt.Print("\n")
    fmt.Println("OR result")
    fmt.Println("========")
    orResult := value1 | value2 // 11111111 [255]
    fmt.Printf("%08b | %08b = %08b\n", value1, value2, orResult)

    // NOT 비트 연산
    fmt.Print("\n")
    fmt.Println("NOT result")
    fmt.Println("========")
    notResult1 := ^value1 // 01000100 [68]
    notResult2 := ^value2 // 00100001 [33]
    fmt.Printf("^%08b: %08b\n", value1, notResult1)
    fmt.Printf("^%08b: %08b\n", value2, notResult2)

    // XOR 비트 연산
    fmt.Print("\n")
    fmt.Println("XOR result")
    fmt.Println("========")
    xorResult := value1 ^ value2 // 01100101 [101]
    fmt.Printf("%08b ^ %08b = %08b\n", value1, value2, xorResult)

    // AND NOT 비트 연산
    fmt.Print("\n")
    fmt.Println("AND NOT result")
    fmt.Println("========")
    andNotResult := value1 &^ value2 // 00100001 [33]
    fmt.Printf("%08b &^ %08b = %08b\n", value1, value2, andNotResult)

    // 좌측 시프트 연산
    fmt.Print("\n")
    fmt.Println("Left shift result")
    fmt.Println("========")
    leftShiftResult := shiftValue << shiftLength // 01001100 [76]
    fmt.Printf("%08b << %08b = %08b\n", shiftValue, shiftLength, leftShiftResult)

    // 우측 시프트 연산
    fmt.Print("\n")
    fmt.Println("Right shift result")
    fmt.Println("========")
    rightShiftResult := shiftValue >> shiftLength // 00000100 [4]
    fmt.Printf("%08b >> %08b = %08b\n", shiftValue, shiftLength, rightShiftResult)
}
