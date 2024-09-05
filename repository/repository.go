package repository

import (
	"github.com/yash-sojitra/todo/internal/models"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (r *TaskRepository) CreateTask(task *models.Task) error {
	return r.db.Create(&task).Error
}

func (r *TaskRepository) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) GetTaskByID(ID int) (*models.Task, error) {
	var task models.Task
	err := r.db.First(&task, ID).Error
	return &task, err
}

func (r *TaskRepository) UpdateTask(task *models.Task) error {
	return r.db.Save(task).Error
}

func (r *TaskRepository) DeleteTask(ID int) error {
	return r.db.Delete(&models.Task{}, ID).Error
}
