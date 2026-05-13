package repository

import (
	"gorm.io/gorm"
	"github.com/bo-patrol/internal/domain"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(task *domain.PatrolTask) error {
	return r.db.Create(task).Error
}

func (r *TaskRepository) GetByID(id uint) (*domain.PatrolTask, error) {
	var task domain.PatrolTask
	err := r.db.Preload("Route").Preload("Assignee").First(&task, id).Error
	return &task, err
}

func (r *TaskRepository) List(page, pageSize int, filters map[string]interface{}) ([]domain.PatrolTask, int64, error) {
	var tasks []domain.PatrolTask
	var total int64

	query := r.db.Model(&domain.PatrolTask{})
	for k, v := range filters {
		query = query.Where(k+" = ?", v)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Preload("Route").Preload("Assignee").
		Offset((page - 1) * pageSize).Limit(pageSize).
		Order("created_at DESC").Find(&tasks).Error; err != nil {
		return nil, 0, err
	}

	return tasks, total, nil
}

func (r *TaskRepository) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&domain.PatrolTask{}).Where("id = ?", id).Updates(updates).Error
}

func (r *TaskRepository) Delete(id uint) error {
	return r.db.Delete(&domain.PatrolTask{}, id).Error
}

func (r *TaskRepository) StartTask(id uint, startTime interface{}) error {
	return r.db.Model(&domain.PatrolTask{}).Where("id = ?", id).Update("status", domain.TaskStatusRunning).Update("start_time", startTime).Error
}

func (r *TaskRepository) CompleteTask(id uint, endTime interface{}) error {
	return r.db.Model(&domain.PatrolTask{}).Where("id = ?", id).Update("status", domain.TaskStatusCompleted).Update("end_time", endTime).Error
}
