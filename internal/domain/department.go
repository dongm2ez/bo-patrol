package domain

import (
	"time"
)

type DepartmentType string

const (
	DeptTypeWater      DepartmentType = "water"
	DeptTypeElectric   DepartmentType = "electric"
	DeptTypeElevator   DepartmentType = "elevator"
	DeptTypeBAS        DepartmentType = "bas"
	DeptTypeAudio      DepartmentType = "audio"
	DeptTypeCivil      DepartmentType = "civil"
	DeptTypeFinance    DepartmentType = "finance"
	DeptTypeAdmin      DepartmentType = "admin"
	DeptTypeOther      DepartmentType = "other"
)

type Department struct {
	ID          uint           `gorm:"primaryKey"`
	Name        string         `gorm:"size:100;not null;uniqueIndex"`
	Code        string         `gorm:"size:50;uniqueIndex"`
	Type        DepartmentType `gorm:"size:50;default:'other'"`
	Description string         `gorm:"type:text"`
	ParentID    *uint          `gorm:"index"`
	Parent      *Department    `gorm:"foreignKey:ParentID"`
	SortOrder   int            `gorm:"default:0"`
	ManagerID   *uint          `gorm:"index"`
	Manager     *User          `gorm:"foreignKey:ManagerID"`
	Phone       string         `gorm:"size:50"`
	Email       string         `gorm:"size:100"`
	Status      int            `gorm:"default:1"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CreateDepartmentRequest struct {
	Name        string         `json:"name" binding:"required"`
	Code        string         `json:"code" binding:"required"`
	Type        DepartmentType `json:"type"`
	Description string         `json:"description"`
	ParentID    *uint          `json:"parentId"`
	SortOrder   int            `json:"sortOrder"`
	ManagerID   *uint          `json:"managerId"`
	Phone       string         `json:"phone"`
	Email       string         `json:"email"`
}

type DepartmentResponse struct {
	ID          uint           `json:"id"`
	Name        string         `json:"name"`
	Code        string         `json:"code"`
	Type        DepartmentType `json:"type"`
	Description string         `json:"description"`
	ParentID    *uint          `json:"parentId"`
	ParentName  string         `json:"parentName"`
	SortOrder   int            `json:"sortOrder"`
	ManagerID   *uint          `json:"managerId"`
	ManagerName string         `json:"managerName"`
	Phone       string         `json:"phone"`
	Email       string         `json:"email"`
	Status      int            `json:"status"`
	CreatedAt   time.Time      `json:"createdAt"`
}

type DepartmentTreeResponse struct {
	DepartmentResponse
	Children []DepartmentTreeResponse `json:"children"`
}

func (d *Department) ToResponse() *DepartmentResponse {
	resp := &DepartmentResponse{
		ID:          d.ID,
		Name:        d.Name,
		Code:        d.Code,
		Type:        d.Type,
		Description: d.Description,
		ParentID:    d.ParentID,
		SortOrder:   d.SortOrder,
		ManagerID:   d.ManagerID,
		Phone:       d.Phone,
		Email:       d.Email,
		Status:      d.Status,
		CreatedAt:   d.CreatedAt,
	}

	if d.Parent != nil {
		resp.ParentName = d.Parent.Name
	}

	if d.Manager != nil {
		resp.ManagerName = d.Manager.Name
	}

	return resp
}

var DefaultDepartmentTypes = []DepartmentType{
	DeptTypeWater,
	DeptTypeElectric,
	DeptTypeElevator,
	DeptTypeBAS,
	DeptTypeAudio,
	DeptTypeCivil,
	DeptTypeFinance,
	DeptTypeAdmin,
	DeptTypeOther,
}

var DepartmentTypeNames = map[DepartmentType]string{
	DeptTypeWater:    "水科",
	DeptTypeElectric: "强电科",
	DeptTypeElevator: "电梯科",
	DeptTypeBAS:      "楼宇自动化科",
	DeptTypeAudio:    "音响设备科",
	DeptTypeCivil:    "土木工程改造科",
	DeptTypeFinance:  "财务科",
	DeptTypeAdmin:    "总务科",
	DeptTypeOther:    "其他",
}
