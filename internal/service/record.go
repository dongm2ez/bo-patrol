package service

import (
	"errors"
	"time"

	"github.com/bo-patrol/internal/domain"
	"github.com/bo-patrol/internal/repository"
)

type RecordService struct {
	recordRepo   *repository.RecordRepository
	routeRepo    *repository.RouteRepository
	museumService *MuseumService
}

func NewRecordService(recordRepo *repository.RecordRepository, routeRepo *repository.RouteRepository, museumService *MuseumService) *RecordService {
	return &RecordService{recordRepo: recordRepo, routeRepo: routeRepo, museumService: museumService}
}

func (s *RecordService) CreateRecord(req *domain.CreateRecordRequest, inspectorID uint) (*domain.PatrolRecord, error) {
	point, err := s.routeRepo.GetPointByQRCode(req.QRCode)
	if err != nil {
		return nil, errors.New("无效的二维码")
	}

	if point.ID != req.PointID {
		return nil, errors.New("二维码与点位不匹配")
	}

	status := domain.RecordStatusNormal
	if req.IsAbnormal {
		status = domain.RecordStatusAbnormal
	}

	record := &domain.PatrolRecord{
		TaskID:       req.TaskID,
		PointID:      req.PointID,
		InspectorID:  inspectorID,
		Status:       status,
		CheckTime:    time.Now(),
		CheckMethod:  domain.CheckMethodQRCode,
		CheckResult:  req.CheckResult,
		Images:       req.Images,
		AbnormalDesc: req.AbnormalDesc,
	}

	if err := s.recordRepo.Create(record); err != nil {
		return nil, err
	}

	if s.museumService != nil {
		if point.CheckEnv && req.Environment != nil && point.GalleryID != nil {
			if err := s.museumService.RecordEnvironment(*point.GalleryID, point.ID, record.ID, req.Environment, inspectorID); err != nil {
				return nil, err
			}
		}

		if point.CheckExhibit && len(req.ExhibitChecks) > 0 {
			for _, exhibitCheck := range req.ExhibitChecks {
				if err := s.museumService.RecordExhibitCheck(exhibitCheck.ExhibitID, record.ID, &exhibitCheck, inspectorID); err != nil {
					return nil, err
				}
			}
		}

		if point.CheckSecurity && req.Security != nil {
			if err := s.museumService.RecordSecurityCheck(point.ID, record.ID, req.Security, inspectorID); err != nil {
				return nil, err
			}
		}
	}

	return record, nil
}

func (s *RecordService) GetRecordByID(id uint) (*domain.RecordResponse, error) {
	record, err := s.recordRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return record.ToResponse(), nil
}

func (s *RecordService) ListRecords(page, pageSize int, filters map[string]interface{}) ([]*domain.RecordResponse, int64, error) {
	records, total, err := s.recordRepo.List(page, pageSize, filters)
	if err != nil {
		return nil, 0, err
	}

	result := make([]*domain.RecordResponse, len(records))
	for i, record := range records {
		result[i] = record.ToResponse()
	}

	return result, total, nil
}
