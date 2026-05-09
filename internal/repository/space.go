package repository

import (
	"gorm.io/gorm"
	"github.com/bo-patrol/internal/domain"
)

type SpaceRepository struct {
	db *gorm.DB
}

func NewSpaceRepository(db *gorm.DB) *SpaceRepository {
	return &SpaceRepository{db: db}
}

func (r *SpaceRepository) CreateLocation(location *domain.SpaceLocation) error {
	return r.db.Create(location).Error
}

func (r *SpaceRepository) GetLocationByID(id uint) (*domain.SpaceLocation, error) {
	var location domain.SpaceLocation
	err := r.db.Preload("Parent").Preload("Gallery").First(&location, id).Error
	return &location, err
}

func (r *SpaceRepository) GetLocationByCode(code string) (*domain.SpaceLocation, error) {
	var location domain.SpaceLocation
	err := r.db.Preload("Parent").Preload("Gallery").Where("code = ?", code).First(&location).Error
	return &location, err
}

func (r *SpaceRepository) GetAllLocations() ([]domain.SpaceLocation, error) {
	var locations []domain.SpaceLocation
	err := r.db.Preload("Parent").Preload("Gallery").Order("sort_order ASC, id ASC").Find(&locations).Error
	return locations, err
}

func (r *SpaceRepository) GetLocationsByType(spaceType domain.SpaceType) ([]domain.SpaceLocation, error) {
	var locations []domain.SpaceLocation
	err := r.db.Preload("Parent").Preload("Gallery").Where("type = ?", spaceType).Order("sort_order ASC, id ASC").Find(&locations).Error
	return locations, err
}

func (r *SpaceRepository) GetLocationsByParent(parentID uint) ([]domain.SpaceLocation, error) {
	var locations []domain.SpaceLocation
	err := r.db.Preload("Parent").Preload("Gallery").Where("parent_id = ?", parentID).Order("sort_order ASC, id ASC").Find(&locations).Error
	return locations, err
}

func (r *SpaceRepository) GetLocationsByFloor(floor int) ([]domain.SpaceLocation, error) {
	var locations []domain.SpaceLocation
	err := r.db.Preload("Parent").Preload("Gallery").Where("floor = ?", floor).Order("sort_order ASC, id ASC").Find(&locations).Error
	return locations, err
}

func (r *SpaceRepository) UpdateLocation(location *domain.SpaceLocation) error {
	return r.db.Save(location).Error
}

func (r *SpaceRepository) DeleteLocation(id uint) error {
	return r.db.Delete(&domain.SpaceLocation{}, id).Error
}

func (r *SpaceRepository) CreateMapElement(element *domain.MapElement) error {
	return r.db.Create(element).Error
}

func (r *SpaceRepository) GetMapElementByID(id uint) (*domain.MapElement, error) {
	var element domain.MapElement
	err := r.db.Preload("Location").First(&element, id).Error
	return &element, err
}

func (r *SpaceRepository) GetMapElementsByLocation(locationID uint) ([]domain.MapElement, error) {
	var elements []domain.MapElement
	err := r.db.Preload("Location").Where("location_id = ?", locationID).Order("sort_order ASC, id ASC").Find(&elements).Error
	return elements, err
}

func (r *SpaceRepository) GetMapElementsByType(elementType string) ([]domain.MapElement, error) {
	var elements []domain.MapElement
	err := r.db.Preload("Location").Where("element_type = ?", elementType).Order("sort_order ASC, id ASC").Find(&elements).Error
	return elements, err
}

func (r *SpaceRepository) UpdateMapElement(element *domain.MapElement) error {
	return r.db.Save(element).Error
}

func (r *SpaceRepository) DeleteMapElement(id uint) error {
	return r.db.Delete(&domain.MapElement{}, id).Error
}

func (r *SpaceRepository) BatchUpdateMapElements(elements []domain.MapElement) error {
	tx := r.db.Begin()
	for i := range elements {
		if err := tx.Save(&elements[i]).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}
