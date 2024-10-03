package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const fileName = "todo.txt"

func WriteToFile() {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()
	fmt.Println("Write text you want to add:-")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error loading input", err)
	}
	// _, err = fmt.Fprintln(file, text)
	// if err != nil {
	// 	fmt.Printf("Error Writing to file: %v\n", err)
	// }
	_, err = file.WriteString(text + "\n")
	if err != nil {
		fmt.Printf("Error Writing to file: %v\n", err)
	}
}
func ReadFromFile() {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error Reading file", err)
	}
	todos := strings.Split(string(data), "\n")
	for n, todo := range todos {
		fmt.Printf("%d -> %s", n, todo)
	}
}

func main() {

	fmt.Println("*******MENU **********")
	// fmt.Println("1.Add a Todo")
	// fmt.Println("Write text you want to add:-")
	// reader := bufio.NewReader(os.Stdin)
	// text, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("Error loading input", err)
	// }
	// fmt.Println(text)
	// WriteToFile()
	ReadFromFile()
}
