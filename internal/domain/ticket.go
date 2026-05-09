package domain

import (
	"time"
)

type TicketType string
type TicketPriority string
type TicketStatus string

const (
	TicketTypeEquipment TicketType = "equipment"
	TicketTypeSafety    TicketType = "safety"
	TicketTypeEnv       TicketType = "environment"
	TicketTypeOther     TicketType = "other"
)

const (
	TicketPriorityUrgent TicketPriority = "urgent"
	TicketPriorityHigh   TicketPriority = "high"
	TicketPriorityMedium TicketPriority = "medium"
	TicketPriorityLow    TicketPriority = "low"
)

const (
	TicketStatusPending   TicketStatus = "pending"
	TicketStatusProcessing TicketStatus = "processing"
	TicketStatusVerifying TicketStatus = "verifying"
	TicketStatusCompleted TicketStatus = "completed"
)

type Ticket struct {
	ID          uint           `gorm:"primaryKey"`
	Title       string         `gorm:"size:200;not null"`
	Type        TicketType     `gorm:"size:20;default:'other'"`
	Priority    TicketPriority `gorm:"size:20;default:'medium'"`
	Status      TicketStatus   `gorm:"size:20;default:'pending'"`
	ReporterID  uint           `gorm:"not null;index"`
	AssigneeID  *uint          `gorm:"index"`
	Description string         `gorm:"type:text"`
	Images      string         `gorm:"type:text"`
	Location    string         `gorm:"size:200"`
	DueTime     *time.Time
	RecordID    *uint
	Reporter    User           `gorm:"foreignKey:ReporterID"`
	Assignee    *User          `gorm:"foreignKey:AssigneeID"`
	Record      *PatrolRecord  `gorm:"foreignKey:RecordID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CreateTicketRequest struct {
	Title       string `json:"title" binding:"required"`
	Type        string `json:"type"`
	Priority    string `json:"priority"`
	Description string `json:"description" binding:"required"`
	Images      string `json:"images"`
	Location    string `json:"location"`
	RecordID    uint   `json:"recordId"`
}

type AssignTicketRequest struct {
	AssigneeID uint `json:"assigneeId" binding:"required"`
}

type TicketResponse struct {
	ID           uint           `json:"id"`
	Title        string         `json:"title"`
	Type         TicketType     `json:"type"`
	Priority     TicketPriority `json:"priority"`
	Status       TicketStatus   `json:"status"`
	ReporterID   uint           `json:"reporterId"`
	ReporterName string         `json:"reporterName"`
	AssigneeID   *uint          `json:"assigneeId"`
	AssigneeName string         `json:"assigneeName"`
	Description  string         `json:"description"`
	Images       string         `json:"images"`
	Location     string         `json:"location"`
	DueTime      *time.Time     `json:"dueTime"`
	RecordID     *uint          `json:"recordId"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
}

func (t *Ticket) ToResponse() *TicketResponse {
	resp := &TicketResponse{
		ID:          t.ID,
		Title:       t.Title,
		Type:        t.Type,
		Priority:    t.Priority,
		Status:      t.Status,
		ReporterID:  t.ReporterID,
		AssigneeID:  t.AssigneeID,
		Description: t.Description,
		Images:      t.Images,
		Location:    t.Location,
		DueTime:     t.DueTime,
		RecordID:    t.RecordID,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
	if t.Reporter.Name != "" {
		resp.ReporterName = t.Reporter.Name
	}
	if t.Assignee != nil {
		resp.AssigneeName = t.Assignee.Name
	}
	return resp
}
