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
		fmt.Println("\n=============================")
		fmt.Println("       TO-DO LIST MENU")
		fmt.Println("=============================")
		fmt.Println("1. View Tasks")
		fmt.Println("2. Add Task")
		fmt.Println("3. Remove Task")
		fmt.Println("4. Exit")
		fmt.Println("=============================")
		fmt.Print("Enter your choice: ")

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
			fmt.Println("\nThank you for using the To-Do List App. Goodbye!")
			return
		default:
			fmt.Println("\nInvalid option. Please try again.")
		}
	}
}

func viewTasks(tasks []string) {
	fmt.Println("\n=============================")
	fmt.Println("          YOUR TASKS")
	fmt.Println("=============================")
	if len(tasks) == 0 {
		fmt.Println("No tasks found. Your to-do list is empty.")
	} else {
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task)
		}
	}
	fmt.Println("=============================")
}

func addTask(tasks []string, scanner *bufio.Scanner) []string {
	fmt.Print("\nEnter the task to add: ")
	if scanner.Scan() {
		task := strings.TrimSpace(scanner.Text())
		if task != "" {
			tasks = append(tasks, task)
			fmt.Println("\nTask added successfully.")
		} else {
			fmt.Println("\nTask cannot be empty.")
		}
	}
	return tasks
}

func removeTask(tasks []string, scanner *bufio.Scanner) []string {
	if len(tasks) == 0 {
		fmt.Println("\nNo tasks to remove. Your to-do list is empty.")
		return tasks
	}

	fmt.Print("\nEnter the number of the task to remove: ")
	if scanner.Scan() {
		var index int
		_, err := fmt.Sscanf(scanner.Text(), "%d", &index)
		if err != nil || index < 1 || index > len(tasks) {
			fmt.Println("\nInvalid task number. Please try again.")
		} else {
			task := tasks[index-1]
			tasks = append(tasks[:index-1], tasks[index:]...)
			fmt.Printf("\nTask \"%s\" removed successfully.\n", task)
		}
	}
	return tasks
}
