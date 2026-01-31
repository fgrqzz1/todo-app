package handlers

import (
	"github.com/gin-gonic/gin"

	"todo-app/internal/repository"
)

func RegisterTaskRoutes(r *gin.Engine, rep repository.TaskRepository) {
	r.GET("/", Home)

	r.GET("/tasks", ListTasks(rep))

	r.GET("/tasks/:id", GetTaskByID(rep))

	r.POST("/tasks", CreateTask(rep))

	r.PATCH("/task/:id/done", MarkDone(rep))

	r.DELETE("/tasks/:id", DeleteTask(rep))

	// todo: написать update
	// r.PUT("/tasks/:id")
}