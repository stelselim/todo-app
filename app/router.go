package app

import (
	"fmt"
	"net/http"
	"todo-app/handler"
	"todo-app/types"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetTasksById(ctx *gin.Context) {
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

func GetAllTasks(ctx *gin.Context) {
	tasks := handler.GetTasks()
	ctx.JSON(http.StatusOK, tasks)
}

func CreateTask(ctx *gin.Context) {
	var newTaskInput types.TaskCreateInput
	if err := ctx.ShouldBindBodyWithJSON(&newTaskInput); err != nil {
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

func UpdateTasksById(ctx *gin.Context) {
	id := ctx.Param("id")
	var updateTaskInput types.TaskUpdateInput
	if err := ctx.ShouldBindBodyWithJSON(&updateTaskInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println(updateTaskInput)
	handler.UpdateTask(id, updateTaskInput.Description, updateTaskInput.Completed)
	ctx.JSON(http.StatusOK, gin.H{"message": "task updated"})
}

func DeleteTaskById(ctx *gin.Context) {
	id := ctx.Param("id")
	handler.DeleteTask(id)
	ctx.JSON(http.StatusOK, gin.H{"message": "task deleted"})
}

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Server is running"})
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/tasks", GetAllTasks)
	router.GET("/tasks/:id", GetTasksById)
	router.POST("/tasks", CreateTask)
	router.PUT("/tasks/:id", UpdateTasksById)
	router.DELETE("/tasks/:id", DeleteTaskById)
	router.GET("/", Ping)

	return router
}
