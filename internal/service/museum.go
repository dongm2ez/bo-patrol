package service

import (
	"errors"
	"time"

	"gorm.io/gorm"
	"github.com/bo-patrol/internal/domain"
	"github.com/bo-patrol/internal/repository"
)

type MuseumService struct {
	museumRepo *repository.MuseumRepository
}

func NewMuseumService(museumRepo *repository.MuseumRepository) *MuseumService {
	return &MuseumService{museumRepo: museumRepo}
}

func (s *MuseumService) CreateGallery(req *domain.CreateGalleryRequest) (*domain.Gallery, error) {
	existing, err := s.museumRepo.GetGalleryByCode(req.Code)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if existing != nil && existing.ID > 0 {
		return nil, errors.New("展厅代码已存在")
	}
	if existing != nil && existing.ID > 0 {
		return nil, errors.New("展厅代码已存在")
	}

	gallery := &domain.Gallery{
		Name:           req.Name,
		Code:           req.Code,
		Zone:           req.Zone,
		Floor:          req.Floor,
		Description:    req.Description,
		Area:           req.Area,
		TemperatureMin: req.TemperatureMin,
		TemperatureMax: req.TemperatureMax,
		HumidityMin:    req.HumidityMin,
		HumidityMax:    req.HumidityMax,
		Status:         1,
	}

	if err := s.museumRepo.CreateGallery(gallery); err != nil {
		return nil, err
	}
	return gallery, nil
}

func (s *MuseumService) GetGallery(id uint) (*domain.Gallery, error) {
	return s.museumRepo.GetGalleryByID(id)
}

func (s *MuseumService) ListGalleries(zone domain.GalleryZone) ([]domain.Gallery, error) {
	return s.museumRepo.ListGalleries(zone)
}

func (s *MuseumService) UpdateGallery(id uint, req *domain.CreateGalleryRequest) (*domain.Gallery, error) {
	gallery, err := s.museumRepo.GetGalleryByID(id)
	if err != nil {
		return nil, err
	}

	gallery.Name = req.Name
	gallery.Code = req.Code
	gallery.Zone = req.Zone
	gallery.Floor = req.Floor
	gallery.Description = req.Description
	gallery.Area = req.Area
	gallery.TemperatureMin = req.TemperatureMin
	gallery.TemperatureMax = req.TemperatureMax
	gallery.HumidityMin = req.HumidityMin
	gallery.HumidityMax = req.HumidityMax

	if err := s.museumRepo.UpdateGallery(gallery); err != nil {
		return nil, err
	}
	return gallery, nil
}

func (s *MuseumService) DeleteGallery(id uint) error {
	return s.museumRepo.DeleteGallery(id)
}

func (s *MuseumService) CreateExhibit(req *domain.CreateExhibitRequest) (*domain.Exhibit, error) {
	existing, err := s.museumRepo.GetExhibitByCode(req.Code)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if existing != nil && existing.ID > 0 {
		return nil, errors.New("展品代码已存在")
	}
	if existing != nil && existing.ID > 0 {
		return nil, errors.New("展品代码已存在")
	}

	exhibit := &domain.Exhibit{
		Name:         req.Name,
		Code:         req.Code,
		GalleryID:    req.GalleryID,
		Level:        req.Level,
		Category:     req.Category,
		Description:  req.Description,
		Material:     req.Material,
		Era:          req.Era,
		Status:       domain.ExhibitStatusOnDisplay,
		LocationDesc: req.LocationDesc,
	}

	if err := s.museumRepo.CreateExhibit(exhibit); err != nil {
		return nil, err
	}
	return exhibit, nil
}

func (s *MuseumService) GetExhibit(id uint) (*domain.Exhibit, error) {
	return s.museumRepo.GetExhibitByID(id)
}

func (s *MuseumService) ListExhibits(galleryID uint, level domain.ExhibitLevel) ([]domain.Exhibit, error) {
	return s.museumRepo.ListExhibits(galleryID, level)
}

func (s *MuseumService) UpdateExhibit(id uint, req *domain.CreateExhibitRequest) (*domain.Exhibit, error) {
	exhibit, err := s.museumRepo.GetExhibitByID(id)
	if err != nil {
		return nil, err
	}

	exhibit.Name = req.Name
	exhibit.Code = req.Code
	exhibit.GalleryID = req.GalleryID
	exhibit.Level = req.Level
	exhibit.Category = req.Category
	exhibit.Description = req.Description
	exhibit.Material = req.Material
	exhibit.Era = req.Era
	exhibit.LocationDesc = req.LocationDesc

	if err := s.museumRepo.UpdateExhibit(exhibit); err != nil {
		return nil, err
	}
	return exhibit, nil
}

func (s *MuseumService) DeleteExhibit(id uint) error {
	return s.museumRepo.DeleteExhibit(id)
}

func (s *MuseumService) RecordEnvironment(galleryID uint, pointID uint, recordID uint, data *domain.EnvironmentData, inspectorID uint) error {
	envRecord := &domain.EnvironmentRecord{
		GalleryID:   galleryID,
		PointID:     pointID,
		RecordID:    recordID,
		Temperature: data.Temperature,
		Humidity:    data.Humidity,
		Lux:         data.Lux,
		CO2:         data.CO2,
		CheckedBy:   inspectorID,
		CheckedAt:   time.Now(),
		Remark:      data.Remark,
	}
	return s.museumRepo.CreateEnvironmentRecord(envRecord)
}

func (s *MuseumService) RecordExhibitCheck(exhibitID uint, recordID uint, data *domain.ExhibitCheckData, inspectorID uint) error {
	checkRecord := &domain.ExhibitCheckRecord{
		ExhibitID:   exhibitID,
		RecordID:    recordID,
		IsIntact:    data.IsIntact,
		IsInPosition: data.IsInPosition,
		HasDamage:   data.HasDamage,
		DamageType:  data.DamageType,
		CaseSealed:  data.CaseSealed,
		Remark:      data.Remark,
		ImageURLs:   data.ImageURLs,
		CheckedBy:   inspectorID,
		CheckedAt:   time.Now(),
	}
	return s.museumRepo.CreateExhibitCheckRecord(checkRecord)
}

func (s *MuseumService) RecordSecurityCheck(pointID uint, recordID uint, data *domain.SecurityCheckData, inspectorID uint) error {
	securityRecord := &domain.SecurityCheckRecord{
		PointID:        pointID,
		RecordID:       recordID,
		CameraWorking:  data.CameraWorking,
		AlarmWorking:   data.AlarmWorking,
		DoorLocked:     data.DoorLocked,
		WindowLocked:   data.WindowLocked,
		FireEquipment:  data.FireEquipment,
		EmergencyLight: data.EmergencyLight,
		Remark:         data.Remark,
		CheckedBy:      inspectorID,
		CheckedAt:      time.Now(),
	}
	return s.museumRepo.CreateSecurityCheckRecord(securityRecord)
}

func (s *MuseumService) GetEnvironmentHistory(galleryID uint, limit int) ([]domain.EnvironmentRecord, error) {
	return s.museumRepo.GetEnvironmentRecordsByGallery(galleryID, limit)
}

func (s *MuseumService) CheckEnvironmentNormal(gallery *domain.Gallery, temp, humidity float64) (bool, string) {
	if temp < gallery.TemperatureMin || temp > gallery.TemperatureMax {
		return false, "温度超出正常范围"
	}
	if humidity < gallery.HumidityMin || humidity > gallery.HumidityMax {
		return false, "湿度超出正常范围"
	}
	return true, ""
}
