package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	tasks := []string{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nTo-Do List Menu:")
		fmt.Println("1. View Tasks")
		fmt.Println("2. Add Task")
		fmt.Println("3. Remove Task")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		if !scanner.Scan() {
			fmt.Println("Error reading input.")
			break
		}
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			viewTasks(tasks)
		case "2":
			tasks = addTask(tasks, scanner)
		case "3":
			tasks = removeTask(tasks, scanner)
		case "4":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func viewTasks(tasks []string) {
	if len(tasks) == 0 {
		fmt.Println("Your to-do list is empty.")
		return
	}

	fmt.Println("\nYour Tasks:")
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}

func addTask(tasks []string, scanner *bufio.Scanner) []string {
	fmt.Print("Enter the task to add: ")
	if scanner.Scan() {
		task := strings.TrimSpace(scanner.Text())
		if task != "" {
			tasks = append(tasks, task)
			fmt.Println("Task added successfully.")
		} else {
			fmt.Println("Task cannot be empty.")
		}
	}
	return tasks
}

func removeTask(tasks []string, scanner *bufio.Scanner) []string {
	if len(tasks) == 0 {
		fmt.Println("Your to-do list is empty. Nothing to remove.")
		return tasks
	}

	fmt.Print("Enter the number of the task to remove: ")
	if scanner.Scan() {
		var index int
		_, err := fmt.Sscanf(scanner.Text(), "%d", &index)
		if err != nil || index < 1 || index > len(tasks) {
			fmt.Println("Invalid task number.")
		} else {
			tasks = append(tasks[:index-1], tasks[index:]...)
			fmt.Println("Task removed successfully.")
		}
	}
	return tasks
}
