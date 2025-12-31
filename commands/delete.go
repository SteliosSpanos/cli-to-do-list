package commands

import (
	"fmt"
	"os"
	"strconv"

	"github.com/SteliosSpanos/cli-to-do-list/pkg"
)

func Delete(args []string) {
	if len(args) != 1 {
		fmt.Println("Error: Please provide the task ID to delete")
		fmt.Println("Usage: ./main delete <id>")
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

	err = taskList.DeleteTask(id)
	if err != nil {
		fmt.Printf("Error deleting task: %v\n", err)
		os.Exit(1)
	}

	err = pkg.SaveToFile(filename, taskList)
	if err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Task Deleted Successfully!")

}
