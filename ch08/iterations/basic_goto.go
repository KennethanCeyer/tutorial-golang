package main

import "fmt"

func main() {
    var answer, name string

askYesOrNo:
    fmt.Print("혹시 관심 있는 연예인이나 가수가 있나요? (네/아니요): ")
    fmt.Scanf("%s", &answer)

    if answer == "아니요" {
        goto askProgramExit
    }

    fmt.Print("이름이 무엇인가요?: ")
    fmt.Scanf("%s", &name)

askProgramExit:
    if name == "" || name == "아니요" {
        var programExitAnswer string
        fmt.Print("그러면 프로그램을 끝낼까요? (네/아니요): ")
        fmt.Scanf("%s", &programExitAnswer)

        if programExitAnswer == "네" {
            goto exit
        }
        goto askYesOrNo
    }

    fmt.Printf("%s님이군요. 관심사를 알려 주어 고마워요!\n", name)

exit:
}
