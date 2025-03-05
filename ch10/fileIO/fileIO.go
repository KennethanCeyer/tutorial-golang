package main

import (
	"fmt"
	"os"
	"strings"
)


func loadData() (name string, phone string) {
	file, err := os.OpenFile("output.txt", os.O_CREATE|os.O_RDONLY, os.FileMode(0644))
	if err != nil {
		fmt.Println(err)
		return "", ""
	}
	defer file.Close()

	f, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return "", ""
	}

	data := make([]byte, f.Size())
	_, err = file.Read(data)
	if err != nil {
		fmt.Println(err)
		return "", ""
	}

	value := string(data)
	if value == "" {
		return "", ""
	}

	splitValues := strings.SplitN(value, ",", 2)
	if len(splitValues) < 2 {
		return "", ""
	}
	return splitValues[0], splitValues[1]
}

func saveData(name string, phone string) {
	file, err := os.OpenFile("output.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.FileMode(0644))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	value := fmt.Sprintf("%s,%s", name, phone)
	_, err = file.Write([]byte(value))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	name, phone := loadData()

	if name == "" {
		fmt.Print("혹시 성함이 어떻게 되시나요?: ")
		fmt.Scanf("%s\n", &name)
	}

	if phone == "" {
		fmt.Printf("연락처는 어떻게 되세요?: ")
		fmt.Scanf("%s", &phone)
	}

	fmt.Println("================================================")
	fmt.Printf("Name:  %s\n", name)
	fmt.Printf("Phone: %s\n", phone)
	fmt.Println("================================================")

	saveData(name, phone)
}
