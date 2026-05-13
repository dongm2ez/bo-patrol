package repository

import (
	"gorm.io/gorm"
	"github.com/bo-patrol/internal/domain"
)

type RouteRepository struct {
	db *gorm.DB
}

func NewRouteRepository(db *gorm.DB) *RouteRepository {
	return &RouteRepository{db: db}
}

func (r *RouteRepository) Create(route *domain.PatrolRoute) error {
	return r.db.Create(route).Error
}

func (r *RouteRepository) GetByID(id uint) (*domain.PatrolRoute, error) {
	var route domain.PatrolRoute
	err := r.db.Preload("Points").First(&route, id).Error
	return &route, err
}

func (r *RouteRepository) List(page, pageSize int) ([]domain.PatrolRoute, int64, error) {
	var routes []domain.PatrolRoute
	var total int64

	offset := (page - 1) * pageSize
	if err := r.db.Model(&domain.PatrolRoute{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Preload("Points").Offset(offset).Limit(pageSize).Find(&routes).Error; err != nil {
		return nil, 0, err
	}

	return routes, total, nil
}

func (r *RouteRepository) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&domain.PatrolRoute{}).Where("id = ?", id).Updates(updates).Error
}

func (r *RouteRepository) Delete(id uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("route_id = ?", id).Delete(&domain.RoutePoint{}).Error; err != nil {
			return err
		}
		return tx.Delete(&domain.PatrolRoute{}, id).Error
	})
}

func (r *RouteRepository) CreatePoint(point *domain.RoutePoint) error {
	return r.db.Create(point).Error
}

func (r *RouteRepository) GetPointByQRCode(qrCode string) (*domain.RoutePoint, error) {
	var point domain.RoutePoint
	err := r.db.Where("qr_code = ?", qrCode).First(&point).Error
	return &point, err
}

func (r *RouteRepository) GetPointsByRouteID(routeID uint) ([]domain.RoutePoint, error) {
	var points []domain.RoutePoint
	err := r.db.Where("route_id = ?", routeID).Order("sort ASC").Find(&points).Error
	return points, err
}
