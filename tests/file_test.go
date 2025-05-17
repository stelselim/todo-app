package tests

import (
	"os"
	"testing"
	"todo-app/helper"
	"todo-app/types"
)

var (
	age  = uint16(26)
	user = types.User{
		Name:    "selim",
		Surname: "ustel",
		Age:     &age,
	}
	task  = types.Task{Id: "1", Description: "read a book", Completed: true, User: &user}
	task2 = types.Task{Id: "2", Description: "buy book", Completed: true, User: &user}
)

func TestGetTaskListFromDatabase(t *testing.T) {
	os.Remove("tasks.json")
	// Check Empty
	tasks := helper.GetTaskListFromDatabase()

	if len(tasks) != 0 {
		t.Errorf("Error getting empty list")
	}

	singleTaskList := []types.Task{task}
	res, _ := helper.SaveTaskListToDatabase(singleTaskList)
	if !res {
		t.Error("Error Saving Tasks to List")
	}

	// Check Empty
	tasks = helper.GetTaskListFromDatabase()

	if len(tasks) == 0 {
		t.Errorf("Error getting single task list")
	}

	os.Remove("tasks.json")
}

func TestSaveTaskListToDatabase(t *testing.T) {
	os.Remove("tasks.json")
	res, _ := helper.SaveTaskListToDatabase([]types.Task{task, task2})
	if !res {
		t.Error("Error Saving Tasks to List")
	}
	tasks := helper.GetTaskListFromDatabase()
	if len(tasks) != 2 {
		t.Error("Error Saving Items to List")
	}
	os.Remove("tasks.json")
}
