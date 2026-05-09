package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/bo-patrol/internal/pkg/db"
	"github.com/bo-patrol/internal/pkg/response"
	"github.com/bo-patrol/internal/service"
)

var (
	statsService *service.StatsService
)

func initStats() {
	statsService = service.NewStatsService(db.DB)
}

func GetOverviewStats(c *gin.Context) {
	stats, err := statsService.GetOverviewStats()
	if err != nil {
		response.ServerError(c, err)
		return
	}

	response.Success(c, stats)
}

func GetCompletionStats(c *gin.Context) {
	stats, err := statsService.GetCompletionStats()
	if err != nil {
		response.ServerError(c, err)
		return
	}

	response.Success(c, stats)
}

func GetAbnormalStats(c *gin.Context) {
	stats, err := statsService.GetAbnormalStats()
	if err != nil {
		response.ServerError(c, err)
		return
	}

	response.Success(c, stats)
}
