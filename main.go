package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)


func main() {
	fmt.Println("Hello to TODO")

	command := flag.String("command", "no command", "command to run")

	flag.Parse()

	runCommand(*command)
}


func runCommand(command string) {
	switch command {
	case "create-task":
		createTask()
	case "create-category":
		createCategory()
	case "register-user":
		registerUser()
	case "login":
		login()
	default:
		fmt.Println("command is not find", command)
	}
}

func createTask() {
	scanner := bufio.NewScanner(os.Stdin)
	var name, duedate, category string

	fmt.Println("please enter the task title")
	scanner.Scan()

	name = scanner.Text()

	fmt.Println("please enter the task category")
	scanner.Scan()

	category = scanner.Text()


	fmt.Println("please enter the task due date")
	scanner.Scan()

	duedate = scanner.Text()


	fmt.Println("task:", name, category, duedate)
}


func createCategory() {
	scanner := bufio.NewScanner(os.Stdin)
	var title, color string

	fmt.Println("please enter the category  title")
	scanner.Scan()

	title = scanner.Text()

	fmt.Println("please enter the category color")
	scanner.Scan()

	color = scanner.Text()

	fmt.Println("category:", title, color)
}

func registerUser() {

	scanner := bufio.NewScanner(os.Stdin)
	var id, email, password string

	fmt.Println("please enter the user email")
	scanner.Scan()

	email = scanner.Text()


	fmt.Println("please enter the user password")
	scanner.Scan()

	password = scanner.Text()

	id = email

	fmt.Println("user:", id, email, password)
}

func login() {
	scanner := bufio.NewScanner(os.Stdin)
	var email, password string

	fmt.Println("please enter the user email")
	scanner.Scan()

	email = scanner.Text()


	fmt.Println("please enter the user password")
	scanner.Scan()

	password = scanner.Text()

	fmt.Println("user:", email, password)
}

