package handelers

import (
	"net/http"

	"github.com/gin-gonic/gin"

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

	// todo: дописать остальные маршруты
}