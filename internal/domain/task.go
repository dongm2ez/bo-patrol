package domain

import (
	"time"
)

type TaskType string
type TaskStatus string

const (
	TaskTypeDaily   TaskType = "daily"
	TaskTypeSpecial TaskType = "special"
	TaskTypeTemp    TaskType = "temp"
)

const (
	TaskStatusPending   TaskStatus = "pending"
	TaskStatusRunning   TaskStatus = "running"
	TaskStatusCompleted TaskStatus = "completed"
	TaskStatusOverdue   TaskStatus = "overdue"
)

type PatrolTask struct {
	ID          uint       `gorm:"primaryKey"`
	Name        string     `gorm:"size:100;not null"`
	Type        TaskType   `gorm:"size:20;default:'daily'"`
	RouteID     uint       `gorm:"not null"`
	AssigneeID  uint       `gorm:"not null;index"`
	Status      TaskStatus `gorm:"size:20;default:'pending'"`
	PlanTime    time.Time  `gorm:"not null"`
	StartTime   *time.Time
	EndTime     *time.Time
	Description string     `gorm:"type:text"`
	CreatorID   uint       `gorm:"not null"`
	Route       PatrolRoute `gorm:"foreignKey:RouteID"`
	Assignee    User       `gorm:"foreignKey:AssigneeID"`
	Records     []PatrolRecord `gorm:"foreignKey:TaskID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CreateTaskRequest struct {
	Name        string    `json:"name" binding:"required"`
	Type        string    `json:"type"`
	RouteID     uint      `json:"routeId" binding:"required"`
	AssigneeID  uint      `json:"assigneeId" binding:"required"`
	PlanTime    time.Time `json:"planTime" binding:"required"`
	Description string    `json:"description"`
}

type TaskResponse struct {
	ID          uint       `json:"id"`
	Name        string     `json:"name"`
	Type        TaskType   `json:"type"`
	RouteID     uint       `json:"routeId"`
	RouteName   string     `json:"routeName"`
	AssigneeID  uint       `json:"assigneeId"`
	AssigneeName string    `json:"assigneeName"`
	Status      TaskStatus `json:"status"`
	PlanTime    time.Time  `json:"planTime"`
	StartTime   *time.Time `json:"startTime"`
	EndTime     *time.Time `json:"endTime"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"createdAt"`
}

func (t *PatrolTask) ToResponse() *TaskResponse {
	resp := &TaskResponse{
		ID:          t.ID,
		Name:        t.Name,
		Type:        t.Type,
		RouteID:     t.RouteID,
		AssigneeID:  t.AssigneeID,
		Status:      t.Status,
		PlanTime:    t.PlanTime,
		StartTime:   t.StartTime,
		EndTime:     t.EndTime,
		Description: t.Description,
		CreatedAt:   t.CreatedAt,
	}
	if t.Route.Name != "" {
		resp.RouteName = t.Route.Name
	}
	if t.Assignee.Name != "" {
		resp.AssigneeName = t.Assignee.Name
	}
	return resp
}
