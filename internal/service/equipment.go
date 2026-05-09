package service

import (
	"errors"
	"time"

	"github.com/bo-patrol/internal/domain"
	"github.com/bo-patrol/internal/repository"
)

type EquipmentService struct {
	equipmentRepo *repository.EquipmentRepository
}

func NewEquipmentService(equipmentRepo *repository.EquipmentRepository) *EquipmentService {
	return &EquipmentService{equipmentRepo: equipmentRepo}
}

func parseDate(dateStr *string) *time.Time {
	if dateStr == nil || *dateStr == "" {
		return nil
	}
	t, err := time.Parse("2006-01-02", *dateStr)
	if err != nil {
		return nil
	}
	return &t
}

func (s *EquipmentService) CreateEquipment(req *domain.CreateEquipmentRequest) (*domain.Equipment, error) {
	existing, _ := s.equipmentRepo.GetByCode(req.Code)
	if existing != nil && existing.ID > 0 {
		return nil, errors.New("设备代码已存在")
	}

	equipment := &domain.Equipment{
		Name:              req.Name,
		Code:              req.Code,
		Category:          req.Category,
		Brand:             req.Brand,
		Model:             req.Model,
		SerialNumber:      req.SerialNumber,
		DepartmentID:      req.DepartmentID,
		LocationID:        req.LocationID,
		LocationDesc:      req.LocationDesc,
		PurchaseDate:      parseDate(req.PurchaseDate),
		PurchasePrice:     req.PurchasePrice,
		SupplierID:        req.SupplierID,
		MaintenanceVendor: req.MaintenanceVendor,
		MaintenanceStart:  parseDate(req.MaintenanceStart),
		MaintenanceEnd:    parseDate(req.MaintenanceEnd),
		LastMaintenance:   parseDate(req.LastMaintenance),
		NextMaintenance:   parseDate(req.NextMaintenance),
		Status:            domain.EquipStatusNormal,
		Specs:             req.Specs,
		Remark:            req.Remark,
	}

	if err := s.equipmentRepo.Create(equipment); err != nil {
		return nil, err
	}
	return equipment, nil
}

func (s *EquipmentService) GetEquipment(id uint) (*domain.Equipment, error) {
	return s.equipmentRepo.GetByID(id)
}

func (s *EquipmentService) GetAllEquipments(page, pageSize int, filters map[string]interface{}) ([]domain.Equipment, int64, error) {
	return s.equipmentRepo.GetAll(page, pageSize, filters)
}

func (s *EquipmentService) GetEquipmentsByDepartment(departmentID uint) ([]domain.Equipment, error) {
	return s.equipmentRepo.GetByDepartment(departmentID)
}

func (s *EquipmentService) GetEquipmentsByLocation(locationID uint) ([]domain.Equipment, error) {
	return s.equipmentRepo.GetByLocation(locationID)
}

func (s *EquipmentService) UpdateEquipment(id uint, req *domain.CreateEquipmentRequest) (*domain.Equipment, error) {
	equipment, err := s.equipmentRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Code != equipment.Code {
		existing, _ := s.equipmentRepo.GetByCode(req.Code)
		if existing != nil && existing.ID > 0 && existing.ID != id {
			return nil, errors.New("设备代码已存在")
		}
	}

	equipment.Name = req.Name
	equipment.Code = req.Code
	equipment.Category = req.Category
	equipment.Brand = req.Brand
	equipment.Model = req.Model
	equipment.SerialNumber = req.SerialNumber
	equipment.DepartmentID = req.DepartmentID
	equipment.LocationID = req.LocationID
	equipment.LocationDesc = req.LocationDesc
	equipment.PurchaseDate = parseDate(req.PurchaseDate)
	equipment.PurchasePrice = req.PurchasePrice
	equipment.SupplierID = req.SupplierID
	equipment.MaintenanceVendor = req.MaintenanceVendor
	equipment.MaintenanceStart = parseDate(req.MaintenanceStart)
	equipment.MaintenanceEnd = parseDate(req.MaintenanceEnd)
	equipment.LastMaintenance = parseDate(req.LastMaintenance)
	equipment.NextMaintenance = parseDate(req.NextMaintenance)
	equipment.Specs = req.Specs
	equipment.Remark = req.Remark

	if err := s.equipmentRepo.Update(equipment); err != nil {
		return nil, err
	}
	return equipment, nil
}

func (s *EquipmentService) DeleteEquipment(id uint) error {
	return s.equipmentRepo.Delete(id)
}

func (s *EquipmentService) CreateSupplier(req *domain.CreateSupplierRequest) (*domain.Supplier, error) {
	supplier := &domain.Supplier{
		Name:    req.Name,
		Code:    req.Code,
		Contact: req.Contact,
		Phone:   req.Phone,
		Email:   req.Email,
		Address: req.Address,
		Type:    req.Type,
		Status:  1,
	}

	if err := s.equipmentRepo.CreateSupplier(supplier); err != nil {
		return nil, err
	}
	return supplier, nil
}

func (s *EquipmentService) GetAllSuppliers() ([]domain.Supplier, error) {
	return s.equipmentRepo.GetAllSuppliers()
}

func (s *EquipmentService) GetMaintenanceRecords(equipmentID uint, limit int) ([]domain.MaintenanceRecord, error) {
	return s.equipmentRepo.GetMaintenanceRecords(equipmentID, limit)
}

func (s *EquipmentService) CreateMaintenanceRecord(record *domain.MaintenanceRecord) error {
	return s.equipmentRepo.CreateMaintenanceRecord(record)
}
