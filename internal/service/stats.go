package service

import (
	"gorm.io/gorm"
	"github.com/bo-patrol/internal/domain"
)

type StatsService struct {
	db *gorm.DB
}

func NewStatsService(db *gorm.DB) *StatsService {
	return &StatsService{db: db}
}

type OverviewStats struct {
	TotalTasks        int64 `json:"totalTasks"`
	CompletedTasks    int64 `json:"completedTasks"`
	TotalRecords      int64 `json:"totalRecords"`
	AbnormalRecords   int64 `json:"abnormalRecords"`
	TotalTickets      int64 `json:"totalTickets"`
	PendingTickets    int64 `json:"pendingTickets"`
	TotalUsers        int64 `json:"totalUsers"`
	CompletionRate    float64 `json:"completionRate"`
}

func (s *StatsService) GetOverviewStats() (*OverviewStats, error) {
	var stats OverviewStats

	s.db.Model(&domain.PatrolTask{}).Count(&stats.TotalTasks)
	s.db.Model(&domain.PatrolTask{}).Where("status = ?", "completed").Count(&stats.CompletedTasks)
	s.db.Model(&domain.PatrolRecord{}).Count(&stats.TotalRecords)
	s.db.Model(&domain.PatrolRecord{}).Where("status = ?", "abnormal").Count(&stats.AbnormalRecords)
	s.db.Model(&domain.Ticket{}).Count(&stats.TotalTickets)
	s.db.Model(&domain.Ticket{}).Where("status IN ?", []string{"pending", "processing"}).Count(&stats.PendingTickets)
	s.db.Model(&domain.User{}).Count(&stats.TotalUsers)

	if stats.TotalTasks > 0 {
		stats.CompletionRate = float64(stats.CompletedTasks) / float64(stats.TotalTasks) * 100
	}

	return &stats, nil
}

type DailyStats struct {
	Date     string `json:"date"`
	Count    int64  `json:"count"`
	Complete int64  `json:"complete"`
}

func (s *StatsService) GetCompletionStats() ([]DailyStats, error) {
	var stats []DailyStats
	
	err := s.db.Raw(`
		SELECT 
			DATE(plan_time) as date, 
			COUNT(*) as count,
			SUM(CASE WHEN status = 'completed' THEN 1 ELSE 0 END) as complete
		FROM patrol_tasks 
		GROUP BY DATE(plan_time) 
		ORDER BY date DESC 
		LIMIT 7
	`).Scan(&stats).Error

	return stats, err
}

type AbnormalStats struct {
	Zone      string `json:"zone"`
	AbnormalCount int64 `json:"abnormalCount"`
	TotalCount int64 `json:"totalCount"`
	Rate float64 `json:"rate"`
}

func (s *StatsService) GetAbnormalStats() ([]AbnormalStats, error) {
	var stats []AbnormalStats
	
	err := s.db.Raw(`
		SELECT 
			pt.zone,
			SUM(CASE WHEN pr.status = 'abnormal' THEN 1 ELSE 0 END) as abnormal_count,
			COUNT(*) as total_count
		FROM patrol_records pr
		INNER JOIN route_points rp ON pr.point_id = rp.id
		INNER JOIN patrol_routes pt ON rp.route_id = pt.id
		GROUP BY pt.zone
	`).Scan(&stats).Error

	for i := range stats {
		if stats[i].TotalCount > 0 {
			stats[i].Rate = float64(stats[i].AbnormalCount) / float64(stats[i].TotalCount) * 100
		}
	}

	return stats, err
}
