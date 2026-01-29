package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"todo-app/internal/models"
	"todo-app/internal/storage"
)

func RegisterTaskRoutes(r *gin.Engine, rep storage.GormTaskRepository) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"messange": "Главная страница",
		})
	})

	r.GET("/tasks", func(c *gin.Context) {
		tasks, err := rep.ListTasks(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to list tasks",
			})
			return
		}
		c.JSON(http.StatusOK, tasks)
	})

	r.GET("/tasks/:id", func (c *gin.Context)  {
		idParam := c.Param("id")

		id, err := strconv.ParseUint(idParam, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid task id",
			})
			return
		}

		task, err := rep.GetTaskByID(c.Request.Context(), uint(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Faild to get task",
			})
			return
		}

		c.JSON(http.StatusOK, task)
	})

	r.POST("/tasks", func(c *gin.Context) {
		var input models.Task

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request body",
			})
			return
		}

		if err := rep.CreateTask(c.Request.Context(), &input); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to create task",
			})
		}

		c.JSON(http.StatusCreated, input)
	})

	r.PATCH("/task/:id/done", func(c *gin.Context){
		idParam := c.Param("id")
		
		id, err := strconv.ParseUint(idParam, 10, 64)
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid task id",
			})
			return
		}

		if err := rep.MarkDone(c.Request.Context(), uint(id)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H {
				"error": "Failed to mark task as done",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Task marked as done",
		})
	})

	r.DELETE("/tasks/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		
		id, err := strconv.ParseUint(idParam, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid task id",
			})
			return
		}

		if err := rep.DeleteTask(c.Request.Context(), uint(id)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H {
				"error": "Failed to delete task",
			})
			return
		}

		c.JSON(http.StatusNoContent, gin.H {
			"message": "Task deleted",
		})
	})

	// todo: написать update
	// r.PUT("/tasks/:id")
}