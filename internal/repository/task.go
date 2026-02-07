package repository

import (
	"context"
	"errors"
	"todo-app/internal/models"
)

var ErrTaskNotFound = errors.New("Task not found")

type TaskRepository interface{
	CreateTask(ctx context.Context, task *models.Task) error
	GetTaskByID(ctx context.Context, id uint) (*models.Task, error)
	ListTasks(ctx context.Context) ([]*models.Task, error)
	MarkDone(ctx context.Context, id uint) error
	DeleteTask(ctx context.Context, id uint) error
	UpdateTask(ctx context.Context, id uint, task *models.Task) error
}