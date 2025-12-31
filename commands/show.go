package commands

import (
	"fmt"
	"os"

	"github.com/SteliosSpanos/cli-to-do-list/pkg"
)

func Show() {
	taskList, err := pkg.ReadFromFile(filename)
	if err != nil {
		fmt.Printf("Error reading tasks: %v\n", err)
		os.Exit(1)
	}

	taskList.ShowTasks()
}
