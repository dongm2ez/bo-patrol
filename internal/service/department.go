package service

import (
	"errors"

	"gorm.io/gorm"
	"github.com/bo-patrol/internal/domain"
	"github.com/bo-patrol/internal/repository"
)

type DepartmentService struct {
	deptRepo *repository.DepartmentRepository
}

func NewDepartmentService(deptRepo *repository.DepartmentRepository) *DepartmentService {
	return &DepartmentService{deptRepo: deptRepo}
}

func (s *DepartmentService) CreateDepartment(req *domain.CreateDepartmentRequest) (*domain.Department, error) {
	existing, err := s.deptRepo.GetByCode(req.Code)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if existing != nil && existing.ID > 0 {
		return nil, errors.New("科室代码已存在")
	}
	if existing != nil && existing.ID > 0 {
		return nil, errors.New("科室代码已存在")
	}

	dept := &domain.Department{
		Name:        req.Name,
		Code:        req.Code,
		Type:        domain.DepartmentType(req.Type),
		Description: req.Description,
		ParentID:    req.ParentID,
		SortOrder:   req.SortOrder,
		ManagerID:   req.ManagerID,
		Phone:       req.Phone,
		Email:       req.Email,
		Status:      1,
	}

	if err := s.deptRepo.Create(dept); err != nil {
		return nil, err
	}
	return dept, nil
}

func (s *DepartmentService) GetDepartment(id uint) (*domain.Department, error) {
	return s.deptRepo.GetByID(id)
}

func (s *DepartmentService) GetAllDepartments() ([]domain.Department, error) {
	return s.deptRepo.GetAll()
}

func (s *DepartmentService) GetDepartmentsByType(deptType domain.DepartmentType) ([]domain.Department, error) {
	return s.deptRepo.GetByType(deptType)
}

func (s *DepartmentService) GetDepartmentTree() ([]domain.DepartmentTreeResponse, error) {
	depts, err := s.deptRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return s.buildTree(nil, depts), nil
}

func (s *DepartmentService) buildTree(parentID *uint, depts []domain.Department) []domain.DepartmentTreeResponse {
	var result []domain.DepartmentTreeResponse
	for _, dept := range depts {
		isMatch := false
		if parentID == nil && dept.ParentID == nil {
			isMatch = true
		} else if parentID != nil && dept.ParentID != nil && *parentID == *dept.ParentID {
			isMatch = true
		}

		if isMatch {
			treeNode := domain.DepartmentTreeResponse{
				DepartmentResponse: *dept.ToResponse(),
				Children:           s.buildTree(&dept.ID, depts),
			}
			result = append(result, treeNode)
		}
	}
	return result
}

func (s *DepartmentService) UpdateDepartment(id uint, req *domain.CreateDepartmentRequest) (*domain.Department, error) {
	dept, err := s.deptRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Code != dept.Code {
		existing, err := s.deptRepo.GetByCode(req.Code)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		if existing != nil && existing.ID > 0 && existing.ID != id {
			return nil, errors.New("科室代码已存在")
		}
	}

	dept.Name = req.Name
	dept.Code = req.Code
	dept.Type = domain.DepartmentType(req.Type)
	dept.Description = req.Description
	dept.ParentID = req.ParentID
	dept.SortOrder = req.SortOrder
	dept.ManagerID = req.ManagerID
	dept.Phone = req.Phone
	dept.Email = req.Email

	if err := s.deptRepo.Update(dept); err != nil {
		return nil, err
	}
	return dept, nil
}

func (s *DepartmentService) DeleteDepartment(id uint) error {
	children, err := s.deptRepo.GetByParentID(id)
	if err != nil {
		return err
	}
	if len(children) > 0 {
		return errors.New("该科室下还有子科室，无法删除")
	}
	return s.deptRepo.Delete(id)
}
