package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"todo-app/internal/handelers"
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

	taskRep := storage.NewGormTaskRepository(db)

	router := gin.Default()
	handelers.RegisterTaskRoutes(router, *taskRep)


	router.Run(":8080")
}

// todo: вынести конфиг
