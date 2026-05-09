package service

import (
	"github.com/bo-patrol/internal/domain"
	"github.com/bo-patrol/internal/repository"
)

type TicketService struct {
	ticketRepo *repository.TicketRepository
}

func NewTicketService(ticketRepo *repository.TicketRepository) *TicketService {
	return &TicketService{ticketRepo: ticketRepo}
}

func (s *TicketService) CreateTicket(req *domain.CreateTicketRequest, reporterID uint) (*domain.Ticket, error) {
	ticket := &domain.Ticket{
		Title:       req.Title,
		Type:        domain.TicketType(req.Type),
		Priority:    domain.TicketPriority(req.Priority),
		Status:      domain.TicketStatusPending,
		ReporterID:  reporterID,
		Description: req.Description,
		Images:      req.Images,
		Location:    req.Location,
	}

	if ticket.Type == "" {
		ticket.Type = domain.TicketTypeOther
	}
	if ticket.Priority == "" {
		ticket.Priority = domain.TicketPriorityMedium
	}

	if req.RecordID > 0 {
		ticket.RecordID = &req.RecordID
	}

	if err := s.ticketRepo.Create(ticket); err != nil {
		return nil, err
	}

	return ticket, nil
}

func (s *TicketService) GetTicketByID(id uint) (*domain.TicketResponse, error) {
	ticket, err := s.ticketRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return ticket.ToResponse(), nil
}

func (s *TicketService) ListTickets(page, pageSize int, filters map[string]interface{}) ([]*domain.TicketResponse, int64, error) {
	tickets, total, err := s.ticketRepo.List(page, pageSize, filters)
	if err != nil {
		return nil, 0, err
	}

	result := make([]*domain.TicketResponse, len(tickets))
	for i, ticket := range tickets {
		result[i] = ticket.ToResponse()
	}

	return result, total, nil
}

func (s *TicketService) AssignTicket(id, assigneeID uint) error {
	return s.ticketRepo.Assign(id, assigneeID)
}

func (s *TicketService) ProcessTicket(id uint) error {
	return s.ticketRepo.Process(id)
}

func (s *TicketService) CompleteTicket(id uint) error {
	return s.ticketRepo.Complete(id)
}
