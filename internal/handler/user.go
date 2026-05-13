package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/bo-patrol/internal/pkg/db"
	"github.com/bo-patrol/internal/pkg/response"
	"github.com/bo-patrol/internal/repository"
	"github.com/bo-patrol/internal/service"
)

var (
	userService *service.UserService
)

func initUser() {
	userRepo := repository.NewUserRepository(db.DB)
	userService = service.NewUserService(userRepo)
}

func getPageInfo(c *gin.Context) (int, int) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")
	
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = 10
	}
	
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	
	return page, pageSize
}

func ListUsers(c *gin.Context) {
	page, pageSize := getPageInfo(c)
	
	users, total, err := userService.List(page, pageSize)
	if err != nil {
		response.ServerError(c, err)
		return
	}

	response.Success(c, gin.H{
		"list":     users,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}

func GetUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ParamError(c, "无效的ID参数")
		return
	}
	
	user, err := userService.GetByID(uint(id))
	if err != nil {
		response.NotFound(c, "用户不存在")
		return
	}

	response.Success(c, user)
}

func UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ParamError(c, "无效的ID参数")
		return
	}
	
	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		response.ParamError(c, err.Error())
		return
	}

	delete(updates, "password")
	delete(updates, "id")

	if err := userService.Update(uint(id), updates); err != nil {
		response.ServerError(c, err)
		return
	}

	response.Success(c, nil)
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.ParamError(c, "无效的ID参数")
		return
	}
	
	if err := userService.Delete(uint(id)); err != nil {
		response.ServerError(c, err)
		return
	}

	response.Success(c, nil)
}
