package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/bo-patrol/internal/domain"
	"github.com/bo-patrol/internal/pkg/db"
	"github.com/bo-patrol/internal/pkg/response"
	"github.com/bo-patrol/internal/repository"
	"github.com/bo-patrol/internal/service"
)

var (
	ticketService *service.TicketService
)

func initTicket() {
	ticketRepo := repository.NewTicketRepository(db.DB)
	ticketService = service.NewTicketService(ticketRepo)
}

func ListTickets(c *gin.Context) {
	page, pageSize := getPageInfo(c)
	
	filters := make(map[string]interface{})
	if status := c.Query("status"); status != "" {
		filters["status"] = status
	}
	if priority := c.Query("priority"); priority != "" {
		filters["priority"] = priority
	}
	if reporterID := c.Query("reporter_id"); reporterID != "" {
		filters["reporter_id"] = reporterID
	}
	if assigneeID := c.Query("assignee_id"); assigneeID != "" {
		filters["assignee_id"] = assigneeID
	}
	
	tickets, total, err := ticketService.ListTickets(page, pageSize, filters)
	if err != nil {
		response.ServerError(c, err)
		return
	}

	response.Success(c, gin.H{
		"list":     tickets,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}

func CreateTicket(c *gin.Context) {
	userID := c.GetUint("user_id")
	
	var req domain.CreateTicketRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, err.Error())
		return
	}

	ticket, err := ticketService.CreateTicket(&req, userID)
	if err != nil {
		response.ServerError(c, err)
		return
	}

	response.Success(c, ticket.ToResponse())
}

func GetTicket(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ParamError(c, "无效的ID参数")
		return
	}
	
	ticket, err := ticketService.GetTicketByID(uint(id))
	if err != nil {
		response.NotFound(c, "工单不存在")
		return
	}

	response.Success(c, ticket)
}

func AssignTicket(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ParamError(c, "无效的ID参数")
		return
	}
	
	var req domain.AssignTicketRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, err.Error())
		return
	}

	if err := ticketService.AssignTicket(uint(id), req.AssigneeID); err != nil {
		response.ServerError(c, err)
		return
	}

	response.Success(c, nil)
}

func ProcessTicket(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ParamError(c, "无效的ID参数")
		return
	}
	
	if err := ticketService.ProcessTicket(uint(id)); err != nil {
		response.ServerError(c, err)
		return
	}

	response.Success(c, nil)
}

func CompleteTicket(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ParamError(c, "无效的ID参数")
		return
	}
	
	if err := ticketService.CompleteTicket(uint(id)); err != nil {
		response.ServerError(c, err)
		return
	}

	response.Success(c, nil)
}
