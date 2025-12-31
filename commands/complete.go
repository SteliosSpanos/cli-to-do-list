package commands

import (
	"fmt"
	"os"
	"strconv"

	"github.com/SteliosSpanos/cli-to-do-list/pkg"
)

func Complete(args []string) {
	if len(args) != 1 {
		fmt.Printf("Error: Please provide the task ID to complete")
		fmt.Printf("Usage: ./main complete <id>")
		os.Exit(1)
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Error ID must be a number: %v\n", err)
		os.Exit(1)
	}

	taskList, err := pkg.ReadFromFile(filename)
	if err != nil {
		fmt.Printf("Error reading tasks: %v\n", err)
		os.Exit(1)
	}

	err = taskList.SetComplete(id)
	if err != nil {
		fmt.Printf("Error completing task: %v\n", err)
		os.Exit(1)
	}

	err = pkg.SaveToFile(filename, taskList)
	if err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Task Completed Successfully!")
}
