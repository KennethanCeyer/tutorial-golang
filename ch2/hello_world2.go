package main

// import 패키지가 여러개라면 아래처럼 괄호를 통해 지정할 수 있다.
// 이때 콤마로 구분하지 않으니 오탈자에 주의하자.
import (
    "fmt"
    // 우리가 다운로드한 패키지는 패키지 이름 그대로 import에 사용한다.
    "github.com/fatih/color"
)

func main() {
    // Printf 함수는 첫번째 인수에 지정된 문자 형식대로
    // 그 다음에 위치한 인수들을 표시시켜주게 된다.
    // %v는 서식 지정자로 인수의 타입에 맞게 문자를 보여준다.
    fmt.Printf("%v %v\n",
        // color 패키지는 우리가 불러온 fatih/color에 속한 패키지이다.
        // 아래 코드를 통해 출력 문자열에 색상을 부여해 줄 수 있다.
        color.RedString("Hello"),
        color.GreenString("World!"))
}

