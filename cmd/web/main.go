package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"todo-app/internal/storage"
)

func main() {
	db, err := storage.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	_ = db // заглушка

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"messange": "Главная страница",
		})
	})

	router.Run(":8080")
}
