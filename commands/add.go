package commands

import (
	"fmt"
	"os"

	"github.com/SteliosSpanos/cli-to-do-list/pkg"
)

const filename = "tasks.json"

func Add(args []string) {
	if len(args) != 1 {
		fmt.Println("Error: Please provide a description for the task")
		fmt.Println("Usage: ./main add <description>")
		os.Exit(1)
	}

	description := args[0]

	taskList, err := pkg.ReadFromFile(filename)
	if err != nil {
		fmt.Printf("Error reading notes: %v\n", err)
		os.Exit(1)
	}

	taskList.AddTask(description)

	err = pkg.SaveToFile(filename, taskList)
	if err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Task Added Successfully!")
}
