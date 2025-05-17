package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"todo-app/app"
	"todo-app/handler"
	"todo-app/types"

	"github.com/stretchr/testify/assert"
)

var (
	newTask = types.Task{Id: "1", Description: "read a book", Completed: true}
)

func TestCreateTask(t *testing.T) {
	router := app.SetupRouter()
	body := `{"description":"wash your hands", "user":{"name":"selim", "surname":"ustel"}}`
	req, _ := http.NewRequest("POST", "/tasks", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "wash your hands")

	// Delete local files.
	os.Remove("tasks.json")
}

func TestGetTasks(t *testing.T) {
	router := app.SetupRouter()
	req, _ := http.NewRequest("GET", "/tasks", nil)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	/// Check Empty Lists
	var emptyTaskList []types.Task
	json.Unmarshal(w.Body.Bytes(), &emptyTaskList)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 0, len(emptyTaskList))

	// Test Single Task
	handler.SaveNewTask(newTask)

	req_2, _ := http.NewRequest("GET", "/tasks", nil)
	req_2.Header.Set("Content-Type", "application/json")

	writer_2 := httptest.NewRecorder()
	router.ServeHTTP(writer_2, req)

	singleTaskList := []types.Task{}
	json.Unmarshal(writer_2.Body.Bytes(), &singleTaskList)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 1, len(singleTaskList))
	assert.Equal(t, "1", singleTaskList[0].Id)
	assert.Equal(t, "read a book", singleTaskList[0].Description)
	assert.Equal(t, true, singleTaskList[0].Completed)
	assert.NotContains(t, singleTaskList, "user")
	assert.NotContains(t, singleTaskList, "name")

	// Delete local files.
	os.Remove("tasks.json")
}

func TestUpdateTask(t *testing.T) {
	handler.SaveNewTask(newTask)

	router := app.SetupRouter()
	body := `{"descriptiom":"buy an apple", "completed":false}`
	req, _ := http.NewRequest("PUT", "/tasks/1", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
	assert.Contains(t, w.Body.String(), "task updated")

	os.Remove("tasks.json")
}

func TestDeleteTask(t *testing.T) {
	handler.SaveNewTask(newTask)

	router := app.SetupRouter()
	req, _ := http.NewRequest("DELETE", "/tasks/1", nil)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "deleted")

	os.Remove("tasks.json")
}
