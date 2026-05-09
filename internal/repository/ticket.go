package repository

import (
	"gorm.io/gorm"
	"github.com/bo-patrol/internal/domain"
)

type TicketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) *TicketRepository {
	return &TicketRepository{db: db}
}

func (r *TicketRepository) Create(ticket *domain.Ticket) error {
	return r.db.Create(ticket).Error
}

func (r *TicketRepository) GetByID(id uint) (*domain.Ticket, error) {
	var ticket domain.Ticket
	err := r.db.Preload("Reporter").Preload("Assignee").First(&ticket, id).Error
	return &ticket, err
}

func (r *TicketRepository) List(page, pageSize int, filters map[string]interface{}) ([]domain.Ticket, int64, error) {
	var tickets []domain.Ticket
	var total int64

	query := r.db.Model(&domain.Ticket{})
	for k, v := range filters {
		query = query.Where(k+" = ?", v)
	}

	query.Count(&total)
	err := query.Preload("Reporter").Preload("Assignee").
		Offset((page - 1) * pageSize).Limit(pageSize).
		Order("created_at DESC").Find(&tickets).Error

	return tickets, total, err
}

func (r *TicketRepository) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&domain.Ticket{}).Where("id = ?", id).Updates(updates).Error
}

func (r *TicketRepository) Assign(id uint, assigneeID uint) error {
	return r.db.Model(&domain.Ticket{}).Where("id = ?", id).Updates(map[string]interface{}{
		"assignee_id": assigneeID,
		"status":      domain.TicketStatusProcessing,
	}).Error
}

func (r *TicketRepository) Process(id uint) error {
	return r.db.Model(&domain.Ticket{}).Where("id = ?", id).Update("status", domain.TicketStatusVerifying).Error
}

func (r *TicketRepository) Complete(id uint) error {
	return r.db.Model(&domain.Ticket{}).Where("id = ?", id).Update("status", domain.TicketStatusCompleted).Error
}
