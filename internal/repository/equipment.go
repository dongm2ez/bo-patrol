package repository

import (
	"gorm.io/gorm"
	"github.com/bo-patrol/internal/domain"
)

type EquipmentRepository struct {
	db *gorm.DB
}

func NewEquipmentRepository(db *gorm.DB) *EquipmentRepository {
	return &EquipmentRepository{db: db}
}

func (r *EquipmentRepository) Create(equipment *domain.Equipment) error {
	return r.db.Create(equipment).Error
}

func (r *EquipmentRepository) GetByID(id uint) (*domain.Equipment, error) {
	var equipment domain.Equipment
	err := r.db.Preload("Department").Preload("Location").Preload("Supplier").First(&equipment, id).Error
	return &equipment, err
}

func (r *EquipmentRepository) GetByCode(code string) (*domain.Equipment, error) {
	var equipment domain.Equipment
	err := r.db.Preload("Department").Preload("Location").Preload("Supplier").Where("code = ?", code).First(&equipment).Error
	return &equipment, err
}

func (r *EquipmentRepository) GetAll(page, pageSize int, filters map[string]interface{}) ([]domain.Equipment, int64, error) {
	var equipments []domain.Equipment
	var total int64

	query := r.db.Model(&domain.Equipment{}).Preload("Department").Preload("Location").Preload("Supplier")

	for k, v := range filters {
		query = query.Where(k+" = ?", v)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if pageSize > 0 {
		offset := (page - 1) * pageSize
		query = query.Offset(offset).Limit(pageSize)
	}

	err := query.Order("created_at DESC").Find(&equipments).Error
	return equipments, total, err
}

func (r *EquipmentRepository) GetByDepartment(departmentID uint) ([]domain.Equipment, error) {
	var equipments []domain.Equipment
	err := r.db.Preload("Department").Preload("Location").Preload("Supplier").Where("department_id = ?", departmentID).Find(&equipments).Error
	return equipments, err
}

func (r *EquipmentRepository) GetByLocation(locationID uint) ([]domain.Equipment, error) {
	var equipments []domain.Equipment
	err := r.db.Preload("Department").Preload("Location").Preload("Supplier").Where("location_id = ?", locationID).Find(&equipments).Error
	return equipments, err
}

func (r *EquipmentRepository) GetByCategory(category domain.EquipmentCategory) ([]domain.Equipment, error) {
	var equipments []domain.Equipment
	err := r.db.Preload("Department").Preload("Location").Preload("Supplier").Where("category = ?", category).Find(&equipments).Error
	return equipments, err
}

func (r *EquipmentRepository) Update(equipment *domain.Equipment) error {
	return r.db.Save(equipment).Error
}

func (r *EquipmentRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Equipment{}, id).Error
}

func (r *EquipmentRepository) CreateSupplier(supplier *domain.Supplier) error {
	return r.db.Create(supplier).Error
}

func (r *EquipmentRepository) GetSupplierByID(id uint) (*domain.Supplier, error) {
	var supplier domain.Supplier
	err := r.db.First(&supplier, id).Error
	return &supplier, err
}

func (r *EquipmentRepository) GetAllSuppliers() ([]domain.Supplier, error) {
	var suppliers []domain.Supplier
	err := r.db.Find(&suppliers).Error
	return suppliers, err
}

func (r *EquipmentRepository) UpdateSupplier(supplier *domain.Supplier) error {
	return r.db.Save(supplier).Error
}

func (r *EquipmentRepository) CreateMaintenanceRecord(record *domain.MaintenanceRecord) error {
	return r.db.Create(record).Error
}

func (r *EquipmentRepository) GetMaintenanceRecords(equipmentID uint, limit int) ([]domain.MaintenanceRecord, error) {
	var records []domain.MaintenanceRecord
	query := r.db.Preload("Equipment").Where("equipment_id = ?", equipmentID).Order("date DESC")
	if limit > 0 {
		query = query.Limit(limit)
	}
	err := query.Find(&records).Error
	return records, err
}
