package handlers

import (
	"net/http"
	"strconv"
	"todo-app/internal/models"
	"todo-app/internal/repository"

	"github.com/gin-gonic/gin"
)

func ListTasks(rep repository.TaskRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		tasks, err := rep.ListTasks(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to list tasks",
			})
			return
		}
		c.JSON(http.StatusOK, tasks)
	}	
}

func GetTaskByID(rep repository.TaskRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
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
	}
}

func CreateTask(rep repository.TaskRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
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
	}
}

func MarkDone(rep repository.TaskRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
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
	}	
}

func DeleteTask(rep repository.TaskRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
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
	}
}