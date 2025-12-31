package main

import (
	"fmt"
	"os"

	"github.com/SteliosSpanos/cli-to-do-list/commands"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: ./main <command> [arguments]")
		fmt.Println("    add         add task in the list")
		fmt.Println("    delete      delete a task from the list")
		fmt.Println("    complete    set a task as completed")
		fmt.Println("    show        display all the tasks in the list")
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "add":
		commands.Add(args)
	case "delete":
		commands.Delete(args)
	case "complete":
		commands.Complete(args)
	case "show":
		commands.Show()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}
