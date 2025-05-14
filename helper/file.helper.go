package helper

import (
	"encoding/json"
	"fmt"
	"os"
	"todo-app/types"
)

const (
	fileName = "tasks.json"
)

func GetTaskListFromDatabase() []types.Task {
	tasks := []types.Task{}

	fileBytes, err := os.ReadFile(fileName)
	if err != nil {
		// Return empty list
		fmt.Println("Error reading", fileName)
		return tasks
	}

	// Empty file, return empty list.
	if string(fileBytes) == "" {
		return tasks
	}

	if err := json.Unmarshal(fileBytes, &tasks); err != nil {
		// Return empty list
		fmt.Println("Error decoding json: ", err)
		return tasks
	}

	return tasks
}

func SaveTaskListToDatabase(tasks []types.Task) (bool, error) {

	bytes, err := json.MarshalIndent(tasks, "", "\t")
	if err != nil {
		return false, err
	}

	if err := os.WriteFile(fileName, bytes, 0644); err != nil {
		return false, err
	}

	return true, nil
}
