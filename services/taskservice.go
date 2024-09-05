package services

import (
	"github.com/yash-sojitra/todo/internal/models"
	"github.com/yash-sojitra/todo/repository"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (h *TaskService) CreateTask(task *models.Task) error {
	return h.repo.CreateTask(task)
}
func (h *TaskService) GetAllTasks() ([]models.Task, error) {
	return h.repo.GetAllTasks()
}
func (h *TaskService) GetTaskByID(ID int) (*models.Task, error) {
	return h.repo.GetTaskByID(ID)
}
func (h *TaskService) UpdateTask(task *models.Task) error {
	return h.repo.UpdateTask(task)
}
func (h *TaskService) DeleteTask(ID int) error {
	return h.repo.DeleteTask(ID)
}