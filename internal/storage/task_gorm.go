package storage

import (
	"context"
	"errors"
	"fmt"
	"todo-app/internal/models"
	"todo-app/internal/repository"

	"gorm.io/gorm"
)

type GormTaskRepository struct {
	db *gorm.DB
}

func NewGormTaskRepository(db *gorm.DB) *GormTaskRepository {
	return &GormTaskRepository{db: db}
}

func (r *GormTaskRepository) CreateTask(ctx context.Context, task *models.Task) error {
	if err := r.db.WithContext(ctx).Create(task).Error; err != nil {
		return fmt.Errorf("Create task error: %w", err)
	}
	return nil
}

func (r *GormTaskRepository) GetTaskByID(ctx context.Context, id uint) (*models.Task, error) {
	var task models.Task
	if err := r.db.WithContext(ctx).First(&task, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository.ErrTaskNotFound
		}
		return nil, fmt.Errorf("Get task by id error: %w", err)
	}
	return &task, nil
}

func (r *GormTaskRepository) ListTasks(ctx context.Context) ([]*models.Task, error) {
	var tasks []*models.Task
	if err := r.db.WithContext(ctx).Find(&tasks).Error; err != nil {
		return nil, fmt.Errorf("List tasks error: %w", err)
	}
	return tasks, nil
}

func (r *GormTaskRepository) MarkDone(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Model(&models.Task{}).Where("id = ?", id).Update("done", true).Error; err != nil {
		return fmt.Errorf("Mark task as done error: %w", err)
	}
	return nil
}

func (r *GormTaskRepository) DeleteTask(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&models.Task{}, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return repository.ErrTaskNotFound
		}
		return fmt.Errorf("Delete task error: %w", err)
	}
	return nil
}

func (r *GormTaskRepository) UpdateTask(ctx context.Context, id uint, task *models.Task) error {
	if err := r.db.WithContext(ctx).Model(&models.Task{}).Where("id = ?", id).Updates(task).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return repository.ErrTaskNotFound
		}
		return fmt.Errorf("Update task error: %w", err)
	}
	return nil
}
