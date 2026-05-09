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
	recordService *service.RecordService
)

func initRecord() {
	recordRepo := repository.NewRecordRepository(db.DB)
	routeRepo := repository.NewRouteRepository(db.DB)
	recordService = service.NewRecordService(recordRepo, routeRepo, museumService)
}

func ListRecords(c *gin.Context) {
	page, pageSize := getPageInfo(c)
	
	filters := make(map[string]interface{})
	if status := c.Query("status"); status != "" {
		filters["status"] = status
	}
	if taskID := c.Query("task_id"); taskID != "" {
		filters["task_id"] = taskID
	}
	if inspectorID := c.Query("inspector_id"); inspectorID != "" {
		filters["inspector_id"] = inspectorID
	}
	
	records, total, err := recordService.ListRecords(page, pageSize, filters)
	if err != nil {
		response.ServerError(c, err)
		return
	}

	response.Success(c, gin.H{
		"list":     records,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}

func CreateRecord(c *gin.Context) {
	userID := c.GetUint("user_id")
	
	var req domain.CreateRecordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, err.Error())
		return
	}

	record, err := recordService.CreateRecord(&req, userID)
	if err != nil {
		response.ServerError(c, err)
		return
	}

	response.Success(c, record.ToResponse())
}

func GetRecord(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	record, err := recordService.GetRecordByID(uint(id))
	if err != nil {
		response.NotFound(c, "记录不存在")
		return
	}

	response.Success(c, record)
}
