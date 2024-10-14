package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const fileName string = "todo.txt"

func openFile() (*os.File, error) {
	// TODO: In case of OpenFile(), CreateFile() should be used
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return file, nil
}
func UpdateTodo() {
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

	fmt.Println("Choose which one to update??")
	var choice int
	fmt.Scanln(&choice)
	fmt.Print("Write text you want :- ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error loading input", err)
	}
	todos[choice] = strings.TrimSpace(text)
	file, err := openFile()
	if err != nil {
		log.Fatal(err)
	}
	file.Truncate(0)
	defer file.Close()
	for _, todo := range todos {
		_, err := file.WriteString(todo + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Item updated successfully")
}
func DeleteTodo() {

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
	fmt.Println("Choose which one to delete:- ")
	var choice int
	fmt.Scanln(&choice)
	var newTodos []string
	for n, todo := range todos {
		if n == choice {
			continue
		}
		newTodos = append(newTodos, todo)
	}
	file, err := openFile()
	if err != nil {
		log.Fatal(err)
	}
	file.Truncate(0)
	defer file.Close()
	for _, newtodo := range newTodos {
		_, err := file.WriteString(newtodo + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Item Deleted Successfully.")

}

func WriteToFile() {
	// file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	// if err != nil {
	// 	fmt.Printf("Error opening file: %v\n", err)
	// 	return
	// }
	file, err := openFile()
	if err != nil {
		log.Fatal(err)
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
func DeleteAll() {
	file, err := openFile()
	if err != nil {
		log.Fatal(err)
	}
	file.Truncate(0)
	defer file.Close()
	fmt.Println("All Todos deleted successfully.")

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
	** 	5. Delete All 	**
	** 	6. Exit 		**`
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
		UpdateTodo()
	case "4":
		DeleteTodo()
	case "5":
		DeleteAll()
	case "6":
		os.Exit(1)
	default:
		return
	}
}
