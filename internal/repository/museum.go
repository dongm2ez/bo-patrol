package repository

import (
	"gorm.io/gorm"
	"github.com/bo-patrol/internal/domain"
)

type MuseumRepository struct {
	db *gorm.DB
}

func NewMuseumRepository(db *gorm.DB) *MuseumRepository {
	return &MuseumRepository{db: db}
}

func (r *MuseumRepository) CreateGallery(gallery *domain.Gallery) error {
	return r.db.Create(gallery).Error
}

func (r *MuseumRepository) GetGalleryByID(id uint) (*domain.Gallery, error) {
	var gallery domain.Gallery
	err := r.db.First(&gallery, id).Error
	return &gallery, err
}

func (r *MuseumRepository) GetGalleryByCode(code string) (*domain.Gallery, error) {
	var gallery domain.Gallery
	err := r.db.Where("code = ?", code).First(&gallery).Error
	return &gallery, err
}

func (r *MuseumRepository) ListGalleries(zone domain.GalleryZone) ([]domain.Gallery, error) {
	var galleries []domain.Gallery
	query := r.db
	if zone != "" {
		query = query.Where("zone = ?", zone)
	}
	err := query.Find(&galleries).Error
	return galleries, err
}

func (r *MuseumRepository) UpdateGallery(gallery *domain.Gallery) error {
	return r.db.Save(gallery).Error
}

func (r *MuseumRepository) DeleteGallery(id uint) error {
	return r.db.Delete(&domain.Gallery{}, id).Error
}

func (r *MuseumRepository) CreateExhibit(exhibit *domain.Exhibit) error {
	return r.db.Create(exhibit).Error
}

func (r *MuseumRepository) GetExhibitByID(id uint) (*domain.Exhibit, error) {
	var exhibit domain.Exhibit
	err := r.db.Preload("Gallery").First(&exhibit, id).Error
	return &exhibit, err
}

func (r *MuseumRepository) GetExhibitByCode(code string) (*domain.Exhibit, error) {
	var exhibit domain.Exhibit
	err := r.db.Preload("Gallery").Where("code = ?", code).First(&exhibit).Error
	return &exhibit, err
}

func (r *MuseumRepository) ListExhibits(galleryID uint, level domain.ExhibitLevel) ([]domain.Exhibit, error) {
	var exhibits []domain.Exhibit
	query := r.db.Preload("Gallery")
	if galleryID > 0 {
		query = query.Where("gallery_id = ?", galleryID)
	}
	if level != "" {
		query = query.Where("level = ?", level)
	}
	err := query.Find(&exhibits).Error
	return exhibits, err
}

func (r *MuseumRepository) UpdateExhibit(exhibit *domain.Exhibit) error {
	return r.db.Save(exhibit).Error
}

func (r *MuseumRepository) DeleteExhibit(id uint) error {
	return r.db.Delete(&domain.Exhibit{}, id).Error
}

func (r *MuseumRepository) CreateEnvironmentRecord(record *domain.EnvironmentRecord) error {
	return r.db.Create(record).Error
}

func (r *MuseumRepository) GetEnvironmentRecordsByGallery(galleryID uint, limit int) ([]domain.EnvironmentRecord, error) {
	var records []domain.EnvironmentRecord
	err := r.db.Preload("Gallery").Where("gallery_id = ?", galleryID).
		Order("checked_at DESC").Limit(limit).Find(&records).Error
	return records, err
}

func (r *MuseumRepository) CreateExhibitCheckRecord(record *domain.ExhibitCheckRecord) error {
	return r.db.Create(record).Error
}

func (r *MuseumRepository) GetExhibitCheckRecords(exhibitID uint, limit int) ([]domain.ExhibitCheckRecord, error) {
	var records []domain.ExhibitCheckRecord
	err := r.db.Preload("Exhibit").Where("exhibit_id = ?", exhibitID).
		Order("checked_at DESC").Limit(limit).Find(&records).Error
	return records, err
}

func (r *MuseumRepository) CreateSecurityCheckRecord(record *domain.SecurityCheckRecord) error {
	return r.db.Create(record).Error
}

func (r *MuseumRepository) GetSecurityCheckRecords(pointID uint, limit int) ([]domain.SecurityCheckRecord, error) {
	var records []domain.SecurityCheckRecord
	err := r.db.Where("point_id = ?", pointID).
		Order("checked_at DESC").Limit(limit).Find(&records).Error
	return records, err
}
