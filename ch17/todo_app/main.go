package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)


type Task struct {
    ID   int
    Name string
}


var (
    tasks       []Task
    taskCounter int
    mutex       sync.Mutex
    wg          sync.WaitGroup
    taskChan    chan Task
)


func init() {
    // 로그 파일 설정
    file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        fmt.Println("로그 파일을 열 수 없습니다:", err)
        os.Exit(1)
    }
    log.SetOutput(file)
    log.Println("프로그램 시작")
}


func main() {
    taskChan = make(chan Task)
    go writeTasksToFile(taskChan)


    // 프로그램 시작 시 파일로부터 할 일 목록 불러오기
    loadTasksFromFile()


    for {
        fmt.Println("할 일 관리 프로그램")
        fmt.Println("1. 할 일 추가")
        fmt.Println("2. 할 일 목록 보기")
        fmt.Println("3. 종료")
        fmt.Print("선택: ")


        var choice int
        fmt.Scanln(&choice)


        switch choice {
        case 1:
            addTask()
        case 2:
            displayTasks()
        case 3:
            close(taskChan)
            wg.Wait()
            fmt.Println("프로그램을 종료합니다.")
            log.Println("프로그램 종료")
            return
        default:
            fmt.Println("올바른 번호를 입력해주세요.")
        }
    }
}


func addTask() {
    fmt.Print("할 일 이름을 입력하세요: ")
    reader := bufio.NewReader(os.Stdin)
    name, _ := reader.ReadString('\n')
    name = strings.TrimSpace(name) // 공백 문자 제거


    mutex.Lock()
    taskCounter++
    task := Task{ID: taskCounter, Name: name}
    tasks = append(tasks, task)
    mutex.Unlock()


    log.Printf("할 일 추가: %v\n", task)


    wg.Add(1)
    taskChan <- task
}


func displayTasks() {
    mutex.Lock()
    defer mutex.Unlock()


    if len(tasks) == 0 {
        fmt.Println("등록된 할 일이 없습니다.")
        return
    }


    fmt.Println("현재 할 일 목록:")
    for _, task := range tasks {
        fmt.Printf("%d. %s\n", task.ID, task.Name)
    }
}


func writeTasksToFile(ch <-chan Task) {
    file, err := os.OpenFile("tasks.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Println("파일 열기 실패:", err)
        return
    }
    defer file.Close()


    for task := range ch {
        _, err := file.WriteString(fmt.Sprintf("%d,%s\n", task.ID, task.Name))
        if err != nil {
            log.Println("파일 쓰기 실패:", err)
        } else {
            log.Printf("파일에 할 일 저장: %v\n", task)
        }
        wg.Done()
    }
}


func loadTasksFromFile() {
    file, err := os.Open("tasks.txt")
    if err != nil {
        if os.IsNotExist(err) {
            log.Println("기존 할 일 파일이 없습니다.")
            return
        }
        log.Println("파일 읽기 실패:", err)
        return
    }
    defer file.Close()


    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.SplitN(line, ",", 2)
        if len(parts) != 2 {
            continue
        }
        id, err := strconv.Atoi(parts[0])
        if err != nil {
            continue
        }
        name := parts[1]
        task := Task{ID: id, Name: name}
        tasks = append(tasks, task)
        if id > taskCounter {
            taskCounter = id
        }
    }


    if err := scanner.Err(); err != nil {
        log.Println("파일 스캔 중 오류 발생:", err)
    }


    log.Printf("총 %d개의 할 일을 불러왔습니다.\n", len(tasks))
}
