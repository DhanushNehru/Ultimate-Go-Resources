package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/yourusername/go-task-manager/pkg/tasks"
)

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  taskmanager add <title> [--due YYYY-MM-DD] [--priority low|medium|high]")
	fmt.Println("  taskmanager list [--all]")
	fmt.Println("  taskmanager complete <taskID>")
	fmt.Println("  taskmanager remove <taskID>")
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Task title is required")
			return
		}
		title := os.Args[2]
		due := ""
		priority := "medium"

		for i := 3; i < len(os.Args); i++ {
			arg := os.Args[i]
			if arg == "--due" && i+1 < len(os.Args) {
				due = os.Args[i+1]
				i++
			} else if arg == "--priority" && i+1 < len(os.Args) {
				priority = os.Args[i+1]
				i++
			}
		}

		task := tasks.NewTask(title, due, priority)
		err := tasks.AddTask(task)
		if err != nil {
			fmt.Println("Error adding task:", err)
			return
		}
		fmt.Println("Task added successfully!")

	case "list":
		showAll := false
		if len(os.Args) > 2 && os.Args[2] == "--all" {
			showAll = true
		}
		taskList, err := tasks.ListTasks(showAll)
		if err != nil {
			fmt.Println("Error listing tasks:", err)
			return
		}
		if len(taskList) == 0 {
			fmt.Println("No tasks found.")
			return
		}
		fmt.Println("ID\tTitle\tStatus\tDue\tPriority")
		for _, t := range taskList {
			fmt.Printf("%d\t%s\t%s\t%s\t%s\n", t.ID, t.Title, t.Status, t.Due, t.Priority)
		}

	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Error: Task ID is required")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}
		err = tasks.CompleteTask(id)
		if err != nil {
			fmt.Println("Error completing task:", err)
			return
		}
		fmt.Println("Task completed successfully!")

	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("Error: Task ID is required")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}
		err = tasks.RemoveTask(id)
		if err != nil {
			fmt.Println("Error removing task:", err)
			return
		}
		fmt.Println("Task removed successfully!")

	default:
		printUsage()
	}
}
