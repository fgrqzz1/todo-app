package repository

import (
	"context"
	"todo-app/internal/models"
)

type TaskRepository interface{
	CreateTask(ctx context.Context, task *models.Task) error
	GetTaskByID(ctx context.Context, id uint) (*models.Task, error)
	ListTasks(ctx context.Context) ([]*models.Task, error)
	MarkDone(ctx context.Context, id uint) error
	DeleteTask(ctx context.Context, id uint) error
	UpdateTask(ctx context.Context, id uint, task *models.Task) error
}