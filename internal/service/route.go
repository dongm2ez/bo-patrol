package service

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/bo-patrol/internal/domain"
	"github.com/bo-patrol/internal/repository"
)

type RouteService struct {
	routeRepo *repository.RouteRepository
}

func NewRouteService(routeRepo *repository.RouteRepository) *RouteService {
	return &RouteService{routeRepo: routeRepo}
}

func generateQRCode() string {
	b := make([]byte, 16)
	rand.Read(b)
	return "PATROL_" + hex.EncodeToString(b)
}

func (s *RouteService) CreateRoute(req *domain.CreateRouteRequest, createdBy uint) (*domain.PatrolRoute, error) {
	route := &domain.PatrolRoute{
		Name:        req.Name,
		Description: req.Description,
		Zone:        req.Zone,
		CreatedBy:   createdBy,
		Status:      1,
	}

	if err := s.routeRepo.Create(route); err != nil {
		return nil, err
	}

	return route, nil
}

func (s *RouteService) GetRouteByID(id uint) (*domain.RouteResponse, error) {
	route, err := s.routeRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return route.ToResponse(), nil
}

func (s *RouteService) ListRoutes(page, pageSize int) ([]*domain.RouteResponse, int64, error) {
	routes, total, err := s.routeRepo.List(page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	result := make([]*domain.RouteResponse, len(routes))
	for i, route := range routes {
		result[i] = route.ToResponse()
	}

	return result, total, nil
}

func (s *RouteService) UpdateRoute(id uint, updates map[string]interface{}) error {
	return s.routeRepo.Update(id, updates)
}

func (s *RouteService) DeleteRoute(id uint) error {
	return s.routeRepo.Delete(id)
}

func (s *RouteService) CreatePoint(routeID uint, req *domain.CreatePointRequest) (*domain.RoutePoint, error) {
	point := &domain.RoutePoint{
		RouteID:    routeID,
		Name:       req.Name,
		QRCode:     generateQRCode(),
		Latitude:   req.Latitude,
		Longitude:  req.Longitude,
		CheckItems: req.CheckItems,
		Sort:       req.Sort,
		Status:     1,
	}

	if err := s.routeRepo.CreatePoint(point); err != nil {
		return nil, err
	}

	return point, nil
}

func (s *RouteService) GetPointsByRouteID(routeID uint) ([]*domain.PointResponse, error) {
	points, err := s.routeRepo.GetPointsByRouteID(routeID)
	if err != nil {
		return nil, err
	}

	result := make([]*domain.PointResponse, len(points))
	for i, point := range points {
		result[i] = point.ToResponse()
	}

	return result, nil
}
