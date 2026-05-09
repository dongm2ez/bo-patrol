package domain

import (
	"time"
)

type UserRole string

const (
	RoleAdmin     UserRole = "admin"
	RoleInspector UserRole = "inspector"
	RoleRepair    UserRole = "repair"
)

type User struct {
	ID           uint        `gorm:"primaryKey"`
	Username     string      `gorm:"size:50;uniqueIndex;not null"`
	Password     string      `gorm:"size:255;not null"`
	Name         string      `gorm:"size:50;not null"`
	Phone        string      `gorm:"size:20"`
	Email        string      `gorm:"size:100"`
	Role         UserRole    `gorm:"size:20;default:'inspector'"`
	DepartmentID *uint       `gorm:"index"`
	Department   *Department `gorm:"foreignKey:DepartmentID"`
	Avatar       string      `gorm:"size:255"`
	Status       int         `gorm:"default:1"`
	LastLoginAt  *time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type RegisterRequest struct {
	Username     string `json:"username" binding:"required,min=3,max=50"`
	Password     string `json:"password" binding:"required,min=6,max=50"`
	Name         string `json:"name" binding:"required,min=2,max=50"`
	Phone        string `json:"phone"`
	Role         string `json:"role"`
	DepartmentID *uint  `json:"departmentId"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}

type UserResponse struct {
	ID             uint        `json:"id"`
	Username       string      `json:"username"`
	Name           string      `json:"name"`
	Phone          string      `json:"phone"`
	Email          string      `json:"email"`
	Role           UserRole    `json:"role"`
	DepartmentID   *uint       `json:"departmentId"`
	DepartmentName string      `json:"departmentName"`
	Avatar         string      `json:"avatar"`
	Status         int         `json:"status"`
}

func (u *User) ToResponse() *UserResponse {
	resp := &UserResponse{
		ID:           u.ID,
		Username:     u.Username,
		Name:         u.Name,
		Phone:        u.Phone,
		Email:        u.Email,
		Role:         u.Role,
		DepartmentID: u.DepartmentID,
		Avatar:       u.Avatar,
		Status:       u.Status,
	}

	if u.Department != nil {
		resp.DepartmentName = u.Department.Name
	}

	return resp
}
