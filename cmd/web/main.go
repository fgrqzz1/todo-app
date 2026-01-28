package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"todo-app/internal/storage"
)

func main() {
	const dsn = "todo.db" // todo: вынести в end

	db, err := storage.OpenDB(dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := storage.Migrate(db); err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"messange": "Главная страница",
		})
	})

	router.Run(":8080")
}
