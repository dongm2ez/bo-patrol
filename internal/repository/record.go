package repository

import (
	"gorm.io/gorm"
	"github.com/bo-patrol/internal/domain"
)

type RecordRepository struct {
	db *gorm.DB
}

func NewRecordRepository(db *gorm.DB) *RecordRepository {
	return &RecordRepository{db: db}
}

func (r *RecordRepository) Create(record *domain.PatrolRecord) error {
	return r.db.Create(record).Error
}

func (r *RecordRepository) GetByID(id uint) (*domain.PatrolRecord, error) {
	var record domain.PatrolRecord
	err := r.db.Preload("Task").Preload("Point").Preload("Inspector").First(&record, id).Error
	return &record, err
}

func (r *RecordRepository) List(page, pageSize int, filters map[string]interface{}) ([]domain.PatrolRecord, int64, error) {
	var records []domain.PatrolRecord
	var total int64

	query := r.db.Model(&domain.PatrolRecord{})
	for k, v := range filters {
		query = query.Where(k+" = ?", v)
	}

	query.Count(&total)
	err := query.Preload("Task").Preload("Point").Preload("Inspector").
		Offset((page - 1) * pageSize).Limit(pageSize).
		Order("check_time DESC").Find(&records).Error

	return records, total, err
}

func (r *RecordRepository) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&domain.PatrolRecord{}).Where("id = ?", id).Updates(updates).Error
}

func (r *RecordRepository) GetByTaskAndPoint(taskID, pointID uint) (*domain.PatrolRecord, error) {
	var record domain.PatrolRecord
	err := r.db.Where("task_id = ? AND point_id = ?", taskID, pointID).First(&record).Error
	return &record, err
}
