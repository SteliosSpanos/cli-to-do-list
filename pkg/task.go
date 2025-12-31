package pkg

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

func SaveToFile(filename string, taskList *TaskList) error {
	data, err := json.MarshalIndent(taskList, "", " ")
	if err != nil {
		return fmt.Errorf("failed to convert to JSON: %w", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write to file %s: %w", filename, err)
	}

	return nil
}

func ReadFromFile(filename string) (*TaskList, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return &TaskList{Tasks: []Task{}}, nil
		}

		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}

	var taskList TaskList

	if len(data) == 0 {
		return &TaskList{Tasks: []Task{}}, nil
	}

	err = json.Unmarshal(data, &taskList)
	if err != nil {
		return nil, fmt.Errorf("failed to convert from JSON: %w", err)
	}

	return &taskList, nil
}

func (tl *TaskList) AddTask(description string) {
	id := 1
	if len(tl.Tasks) > 0 {
		id = tl.Tasks[len(tl.Tasks)-1].Id + 1
	}

	task := Task{
		Id:          id,
		Description: description,
		Status:      "Not Completed",
	}

	tl.Tasks = append(tl.Tasks, task)
}

func (tl *TaskList) DeleteTask(id int) error {
	for i, task := range tl.Tasks {
		if task.Id == id {
			tl.Tasks = append(tl.Tasks[:i], tl.Tasks[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("failed to find task [ID %d] for deletion", id)
}

func (tl *TaskList) ShowTasks() {
	if len(tl.Tasks) == 0 {
		fmt.Println("No Tasks Found")
		return
	}

	fmt.Println("\n=======================================================")
	fmt.Println("                        YOUR TASKS")
	fmt.Println("=======================================================")

	for _, task := range tl.Tasks {
		fmt.Printf("\n[ID %d] %s  |%s|\n", task.Id, task.Description, task.Status)
	}

	fmt.Println("=======================================================")

}

func (tl *TaskList) SetComplete(id int) error {
	for i := range tl.Tasks {
		if tl.Tasks[i].Id == id {
			tl.Tasks[i].Status = "Completed"
			return nil
		}
	}

	return fmt.Errorf("failed to find task [ID %d] for updating", id)
}
