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
	fmt.Print("Write text you want to add:-")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error loading input", err)
	}
	_, err = file.WriteString(strings.TrimSpace(text) + "\n")
	if err != nil {
		fmt.Printf("Error Writing to file: %v\n", err)
	}
}
func ReadFromFile() {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error Reading file", err)
	}

	todos := strings.Split(strings.TrimSpace(string(data)), "\n")
	if todos[0] == "" {
		fmt.Println("No todos please add some todos.")
		return
	}
	for n, todo := range todos {
		if todo == "" {
			continue
		}
		fmt.Printf("%d -> %s\n", n, todo)
	}
}

func main() {

	menu := `     *******Menu*******
	**	1. List todo 	**
	**	2. Add todo  	**
	** 	3. Update todo 	**
	**	4. Delete todo 	**
	** 	5. Exit 		**`
	fmt.Println(menu)
	var choice string
	fmt.Scanln(&choice)
	switch choice {
	case "1":
		fmt.Println("Here are the list of Todos:- ")
		ReadFromFile()
		break
	case "2":
		WriteToFile()
		break
	case "3":
	case "4":
	case "5":
		return
	default:
		return
	}
}
