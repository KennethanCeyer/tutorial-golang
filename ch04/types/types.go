package main

import "fmt"

func main() {
    const pi float64 = 3.1415926535898932
    const name string = "홍길동"

    // 나이는 음수가 나오지 않기 때문에 unsigned int로 사용한다.
    var age uint8 = 28

    // 속도는 음수가 되면 반대 방향으로 진행할 수 있기 때문에
    // 음수 부호를 지원하는 float32 형태로 작성한다.
    var speed float32 = 60.5

    // 한국인일 경우 true(참), 한국인이 아닐경우 false(거짓) 값을 가지는
    // isKorean 변수는 논리(bool) 타입으로 지정한다.
    var isKorean = true

    fmt.Printf("파이는 %v 입니다.", pi)
    fmt.Print("\n========\n")

    // name은 문자열이기 때문에 %s 서식 지정자를 사용하여 출력하고
    // age는 정수이기 때문에 %d 서식 지정자를 통해 출력한다.
    // 둘 모두 %v를 통해서 출력 또한 가능하다.
    fmt.Printf("이름: %s, 나이: %d\n", name, age)

    // 속도는 실수이기 때문에 %f를 통해 출력한다.
    // %v 서식 지정자로 출력하면 더 많은 소수점 자리를 출력시킬 수 있다.
    // %f로 출력할 경우 6자리의 소수점까지 출력시키게 된다.
    fmt.Printf("속도: %f\n", speed)


    // if 키워드는 아직 배우지 않았지만, 참고를 위해 살펴보기를 바란다.
    // isKorean가 true이면 39번째 라인이 실행되고
    // false이면 41번째 라인이 실행된다.
    // 현재 isKorean은 true이기 때문에 39번째 라인이 실행되게 된다.
    if isKorean {
        fmt.Println("한국인 입니다.")
    } else {
        fmt.Println("외국인 입니다.")
    }
}
