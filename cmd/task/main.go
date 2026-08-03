package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/musanii/go-task-manager/internal/task"
)

func main() {
	service := task.NewService()

	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage:")
		fmt.Println(` task add "Task title"`)
		fmt.Println("task list")
		fmt.Println("task complete <id>")
		fmt.Println("task delete <id>")

		return
	}
	command := args[1]

	switch command {
	case "add":
		addTask(service, args[2:])

	case "list":
		listTasks(service)

	default:
		fmt.Printf("Unknown command: %s\n", command)

	}
}

func addTask(service *task.Service, args []string) {
	title := strings.Join(args, " ")

	createdTask, err := service.Create(title)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf(
		"Task %d created: %s\n",
		createdTask.ID,
		createdTask.Title,
	)
}

func listTasks(service *task.Service) {
	tasks := service.List()

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	for _, currentTask := range tasks {
		status := "[ ]"

		if currentTask.Completed {
			status = "[X]"
		}

		fmt.Printf(
			"%d %s %s\n",
			currentTask.ID,
			status,
			currentTask.Title,
		)
	}
}
