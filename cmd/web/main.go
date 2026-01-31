package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"todo-app/internal/config"
	"todo-app/internal/handlers"
	"todo-app/internal/storage"
)

func main() {
	cfg := config.Load()

	db, err := storage.OpenDB(cfg.DSN)
	if err != nil {
		log.Fatal(err)
	}

	if err := storage.Migrate(db); err != nil {
		log.Fatal(err)
	}

	taskRep := storage.NewGormTaskRepository(db)
	
	gin.SetMode(cfg.GinMode)

	router := gin.Default()
	handlers.RegisterTaskRoutes(router, taskRep)

	if err := router.Run(cfg.HTTPPort); err != nil {
		log.Fatalf("Server disconnect: %v", err)
	}
}