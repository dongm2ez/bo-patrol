package repository

import (
	"gorm.io/gorm"
	"github.com/bo-patrol/internal/domain"
)

type DepartmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) *DepartmentRepository {
	return &DepartmentRepository{db: db}
}

func (r *DepartmentRepository) Create(dept *domain.Department) error {
	return r.db.Create(dept).Error
}

func (r *DepartmentRepository) GetByID(id uint) (*domain.Department, error) {
	var dept domain.Department
	err := r.db.Preload("Parent").Preload("Manager").First(&dept, id).Error
	return &dept, err
}

func (r *DepartmentRepository) GetByCode(code string) (*domain.Department, error) {
	var dept domain.Department
	err := r.db.Where("code = ?", code).First(&dept).Error
	return &dept, err
}

func (r *DepartmentRepository) GetAll() ([]domain.Department, error) {
	var depts []domain.Department
	err := r.db.Preload("Parent").Preload("Manager").Order("sort_order ASC, id ASC").Find(&depts).Error
	return depts, err
}

func (r *DepartmentRepository) GetByType(deptType domain.DepartmentType) ([]domain.Department, error) {
	var depts []domain.Department
	err := r.db.Where("type = ?", deptType).Preload("Parent").Preload("Manager").Order("sort_order ASC, id ASC").Find(&depts).Error
	return depts, err
}

func (r *DepartmentRepository) GetByParentID(parentID uint) ([]domain.Department, error) {
	var depts []domain.Department
	err := r.db.Where("parent_id = ?", parentID).Preload("Parent").Preload("Manager").Order("sort_order ASC, id ASC").Find(&depts).Error
	return depts, err
}

func (r *DepartmentRepository) Update(dept *domain.Department) error {
	return r.db.Save(dept).Error
}

func (r *DepartmentRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Department{}, id).Error
}

func (r *DepartmentRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&domain.Department{}).Count(&count).Error
	return count, err
}
