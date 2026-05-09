package service

import (
	"time"

	"github.com/bo-patrol/internal/domain"
	"github.com/bo-patrol/internal/repository"
)

type TaskService struct {
	taskRepo *repository.TaskRepository
}

func NewTaskService(taskRepo *repository.TaskRepository) *TaskService {
	return &TaskService{taskRepo: taskRepo}
}

func (s *TaskService) CreateTask(req *domain.CreateTaskRequest, creatorID uint) (*domain.PatrolTask, error) {
	task := &domain.PatrolTask{
		Name:        req.Name,
		Type:        domain.TaskType(req.Type),
		RouteID:     req.RouteID,
		AssigneeID:  req.AssigneeID,
		PlanTime:    req.PlanTime,
		Description: req.Description,
		Status:      domain.TaskStatusPending,
		CreatorID:   creatorID,
	}

	if task.Type == "" {
		task.Type = domain.TaskTypeDaily
	}

	if err := s.taskRepo.Create(task); err != nil {
		return nil, err
	}

	return task, nil
}

func (s *TaskService) GetTaskByID(id uint) (*domain.TaskResponse, error) {
	task, err := s.taskRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return task.ToResponse(), nil
}

func (s *TaskService) ListTasks(page, pageSize int, filters map[string]interface{}) ([]*domain.TaskResponse, int64, error) {
	tasks, total, err := s.taskRepo.List(page, pageSize, filters)
	if err != nil {
		return nil, 0, err
	}

	result := make([]*domain.TaskResponse, len(tasks))
	for i, task := range tasks {
		result[i] = task.ToResponse()
	}

	return result, total, nil
}

func (s *TaskService) StartTask(id uint) error {
	return s.taskRepo.StartTask(id, time.Now())
}

func (s *TaskService) CompleteTask(id uint) error {
	return s.taskRepo.CompleteTask(id, time.Now())
}
