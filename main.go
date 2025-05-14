package main

import (
	"fmt"
	"todo-app/handler"
	"todo-app/types"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getTasksById(ctx *gin.Context) {
	id := ctx.Param("id")
	tasks := handler.GetTasks()

	for _, item := range tasks {
		if item.Id == id {
			ctx.JSON(http.StatusOK, item)
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"error": "Task not found"})
}

func getAllTasks(ctx *gin.Context) {
	tasks := handler.GetTasks()
	ctx.JSON(http.StatusOK, tasks)
}

func createTask(ctx *gin.Context) {
	var newTaskInput types.TaskCreateInput
	if err := ctx.ShouldBindBodyWithJSON(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	newTask := types.Task{
		Id:          uuid.NewString(),
		Description: newTaskInput.Description,
		Completed:   false,
		User:        newTaskInput.User,
	}
	handler.SaveNewTask(newTask)
	ctx.JSON(http.StatusOK, gin.H{"message": "new task added", "task": newTask})
}

func updateTasksById(ctx *gin.Context) {
	id := ctx.Param("id")
	var updateTaskInput types.TaskUpdateInput
	if err := ctx.ShouldBindBodyWithJSON(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println(updateTaskInput)
	handler.UpdateTask(id, updateTaskInput.Description, updateTaskInput.Completed)
	ctx.JSON(http.StatusOK, gin.H{"message": "task updated"})
}

func deleteTaskById(ctx *gin.Context) {
	id := ctx.Param("id")
	handler.DeleteTask(id)
	ctx.JSON(http.StatusOK, gin.H{"message": "task deleted"})
}

func ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Server is running"})
}

func main() {
	router := gin.Default()
	router.GET("/tasks", getAllTasks)
	router.GET("/tasks/:id", getTasksById)
	router.POST("/tasks", createTask)
	router.PUT("/tasks/:id", updateTasksById)
	router.DELETE("/tasks/:id", deleteTaskById)
	router.GET("/", ping)
	router.Run(":8080")
}
