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
	routeService *service.RouteService
)

func initRoute() {
	routeRepo := repository.NewRouteRepository(db.DB)
	routeService = service.NewRouteService(routeRepo)
}

func ListRoutes(c *gin.Context) {
	page, pageSize := getPageInfo(c)
	
	routes, total, err := routeService.ListRoutes(page, pageSize)
	if err != nil {
		response.ServerError(c, err)
		return
	}

	response.Success(c, gin.H{
		"list":     routes,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}

func CreateRoute(c *gin.Context) {
	userID := c.GetUint("user_id")
	
	var req domain.CreateRouteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, err.Error())
		return
	}

	route, err := routeService.CreateRoute(&req, userID)
	if err != nil {
		response.ServerError(c, err)
		return
	}

	response.Success(c, route.ToResponse())
}

func GetRoute(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	route, err := routeService.GetRouteByID(uint(id))
	if err != nil {
		response.NotFound(c, "路线不存在")
		return
	}

	response.Success(c, route)
}

func UpdateRoute(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		response.ParamError(c, err.Error())
		return
	}

	delete(updates, "id")

	if err := routeService.UpdateRoute(uint(id), updates); err != nil {
		response.ServerError(c, err)
		return
	}

	response.Success(c, nil)
}

func DeleteRoute(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	if err := routeService.DeleteRoute(uint(id)); err != nil {
		response.ServerError(c, err)
		return
	}

	response.Success(c, nil)
}

func ListPoints(c *gin.Context) {
	routeID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	points, err := routeService.GetPointsByRouteID(uint(routeID))
	if err != nil {
		response.ServerError(c, err)
		return
	}

	response.Success(c, points)
}

func CreatePoint(c *gin.Context) {
	routeID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	var req domain.CreatePointRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, err.Error())
		return
	}

	point, err := routeService.CreatePoint(uint(routeID), &req)
	if err != nil {
		response.ServerError(c, err)
		return
	}

	response.Success(c, point.ToResponse())
}
