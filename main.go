package main

import "fmt"

func main() {
	fmt.Println("This is my first golang project doing alone ")
	var choice int
	for true {

		fmt.Println("_____MENU_______")
		fmt.Println("Choose From Above")
		fmt.Println("1.Add a Task")
		fmt.Println("2. Delete a Task")
		fmt.Println("3. Update a Task")
		fmt.Println("4. Exit")
		fmt.Print("Please choose from Above:")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			AddTask("Clean Bathroom")
			break
		case 4:
			return
		default:
			break
		}
	}

}
func AddTask(task string) {
	fmt.Println(task)
}
