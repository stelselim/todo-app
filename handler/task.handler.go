package handler

import (
	"fmt"
	"slices"
	"todo-app/helper"
	"todo-app/types"
)

func SaveNewTask(task types.Task) {
	fmt.Println("Save a Task")
	tasks := GetTasks()
	tasks = append(tasks, task)
	if _, err := helper.SaveTaskListToDatabase(tasks); err != nil {
		fmt.Println("An error occured while saving")
	}

}

func UpdateTask(taskId string, description string, completed bool) {
	fmt.Println("Update a Task")
	tasks := GetTasks()
	for index, val := range tasks {
		if val.Id == taskId {
			tasks[index].Description = description
			tasks[index].Completed = completed
		}
	}

	if _, err := helper.SaveTaskListToDatabase(tasks); err != nil {
		fmt.Println("An error occured while saving")
	}

}

func DeleteTask(taskId string) {
	fmt.Println("Delete a Task")
	tasks := GetTasks()
	tasks = slices.DeleteFunc(tasks, func(item types.Task) bool {
		return item.Id == taskId
	})
	if _, err := helper.SaveTaskListToDatabase(tasks); err != nil {
		fmt.Println("An error occured while saving")
	}

}

func GetTasks() []types.Task {
	tasks := helper.GetTaskListFromDatabase()
	return tasks
}
