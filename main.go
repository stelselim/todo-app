package main

import (
	"fmt"
	"time"
	"todo-app/handler"
	"todo-app/types"

	"github.com/google/uuid"
)

var (
	name             = "Selim"
	surname          = "Ustel"
	age              = uint16(16)
	taskDesciption   = "Wash your clothes"
	taskId           = uuid.New().String()
	taskDesciption_2 = "Brush your teeth"
	taskId_2         = uuid.New().String()
)

func main() {
	fmt.Println("Welcome to ToDo-Go Project.")

	user := types.User{Name: name, Surname: surname, Age: &age}
	task := types.Task{Id: taskId, Description: taskDesciption, Completed: false, User: &user}
	task_2 := types.Task{Id: taskId_2, Description: taskDesciption_2, Completed: false, User: &user}

	handler.SaveNewTask(task)
	handler.SaveNewTask(task_2)
	handler.GetTasks()
	handler.UpdateTask(task_2.Id, "New Description", true)
	taskId_3 := uuid.NewString()
	handler.SaveNewTask(types.Task{Id: taskId_3, Description: "To be deleted", Completed: false})
	time.Sleep(5 * time.Second)
	handler.DeleteTask(taskId_3)

}
