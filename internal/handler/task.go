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
	taskService *service.TaskService
)

func initTask() {
	taskRepo := repository.NewTaskRepository(db.DB)
	taskService = service.NewTaskService(taskRepo)
}

func ListTasks(c *gin.Context) {
	page, pageSize := getPageInfo(c)
	
	filters := make(map[string]interface{})
	if status := c.Query("status"); status != "" {
		filters["status"] = status
	}
	if assigneeID := c.Query("assignee_id"); assigneeID != "" {
		filters["assignee_id"] = assigneeID
	}
	
	tasks, total, err := taskService.ListTasks(page, pageSize, filters)
	if err != nil {
		response.ServerError(c, err)
		return
	}

	response.Success(c, gin.H{
		"list":     tasks,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}

func CreateTask(c *gin.Context) {
	userID := c.GetUint("user_id")
	
	var req domain.CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, err.Error())
		return
	}

	task, err := taskService.CreateTask(&req, userID)
	if err != nil {
		response.ServerError(c, err)
		return
	}

	response.Success(c, task.ToResponse())
}

func GetTask(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	task, err := taskService.GetTaskByID(uint(id))
	if err != nil {
		response.NotFound(c, "任务不存在")
		return
	}

	response.Success(c, task)
}

func StartTask(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	if err := taskService.StartTask(uint(id)); err != nil {
		response.ServerError(c, err)
		return
	}

	response.Success(c, nil)
}

func CompleteTask(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	if err := taskService.CompleteTask(uint(id)); err != nil {
		response.ServerError(c, err)
		return
	}

	response.Success(c, nil)
}
