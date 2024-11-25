package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type User struct {
	ID int
	Name string
	Email string
	Password string
}

type Task struct {
	ID int
	Title string
	DueDate string
	CategoryID int
	IsDone bool
	UserID int
}

type Category struct {
	ID int
	Title string
	Color string
	UserID int
}

var userStore []User 
var taskStore  []Task 
var categoryStore []Category
var authenticatedUser *User

func main() {
	fmt.Println("Hello to TODO")

	command := flag.String("command", "no command", "command to run")

	flag.Parse()

	for {
		runCommand(*command)

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("please enter another command")
		scanner.Scan()

		*command = scanner.Text()

	}


}


func runCommand(command string) {
	if command != "register-user" && command != "exit" && authenticatedUser == nil {
		login()

		if authenticatedUser == nil {
			return
		}
	}
	switch command {
	case "create-task":
		createTask()
	case "create-category":
		createCategory()
	case "register-user":
		registerUser()
	case "list-task":
		listTask()
	case "login":
		login()
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("command is not find", command)
	}
}

func createTask() {
	if (authenticatedUser != nil) {

	}
	scanner := bufio.NewScanner(os.Stdin)
	var title, duedate, category string

	fmt.Println("please enter the task title")
	scanner.Scan()

	title = scanner.Text()

	fmt.Println("please enter the task category id")
	scanner.Scan()

	category = scanner.Text()

	categoryID, err := strconv.Atoi(category)

	if err != nil {
		fmt.Printf("category id is not valid integer, %v\n", err)

		return
	}

	ifFound := false
	for _, c := range categoryStore {
		if c.ID == categoryID  && c.UserID == authenticatedUser.ID{
			ifFound = true
			break
		}
	}

	if !ifFound {
		fmt.Printf("category id is not found\n")

		return
	}

	fmt.Println("please enter the task due date")
	scanner.Scan()

	duedate = scanner.Text()

	task := Task{
		ID        : len(taskStore) + 1,
		Title     : title,
		DueDate   : duedate,
		CategoryID: categoryID,
		IsDone    : false,
		UserID    : authenticatedUser.ID,

	}

	taskStore = append(taskStore, task)


	fmt.Println("task:", title, category, duedate)
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

	c := Category {
		ID: len(categoryStore) + 1,
		Title: title,
		Color: color,
		UserID: authenticatedUser.ID,
	}

	categoryStore = append(categoryStore, c)

	fmt.Println("category:", title, color)
}

func registerUser() {

	scanner := bufio.NewScanner(os.Stdin)
	var name, email, password string

	fmt.Println("please enter the user name")
	scanner.Scan()

	name = scanner.Text()

	fmt.Println("please enter the user email")
	scanner.Scan()

	email = scanner.Text()


	fmt.Println("please enter the user password")
	scanner.Scan()

	password = scanner.Text()


	id := len(userStore) + 1

	fmt.Println("user:", id, email, password)

	user := User {
		ID: rand.Int(),
		Name: name,
		Email: email,
		Password: password,
	}

	userStore = append(userStore, user)
}

func listTask() {
	for _, task := range taskStore {
		if task.UserID == authenticatedUser.ID {
			fmt.Println(task)
		}
	}
}

func login() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("you must login first")

	fmt.Println("please enter the user email")
	scanner.Scan()

	email := scanner.Text()


	fmt.Println("please enter the user password")
	scanner.Scan()

	password := scanner.Text()

	for _, user := range userStore {
		if user.Email == email && user.Password == password {
			authenticatedUser = &user

			break
		}
	}

	if authenticatedUser == nil {
		fmt.Println("email or password is not correct")
	}

	fmt.Println("user:", email, password)
}

