package service

import (
	"errors"

	"gorm.io/gorm"
	"github.com/bo-patrol/internal/domain"
	"github.com/bo-patrol/internal/repository"
)

type SpaceService struct {
	spaceRepo *repository.SpaceRepository
}

func NewSpaceService(spaceRepo *repository.SpaceRepository) *SpaceService {
	return &SpaceService{spaceRepo: spaceRepo}
}

func (s *SpaceService) CreateLocation(req *domain.CreateSpaceRequest) (*domain.SpaceLocation, error) {
	existing, err := s.spaceRepo.GetLocationByCode(req.Code)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if existing != nil && existing.ID > 0 {
		return nil, errors.New("空间代码已存在")
	}

	location := &domain.SpaceLocation{
		Name:        req.Name,
		Code:        req.Code,
		Type:        domain.SpaceType(req.Type),
		ParentID:    req.ParentID,
		GalleryID:   req.GalleryID,
		Floor:       req.Floor,
		Description: req.Description,
		MapData:     req.MapData,
		CoordsX:     req.CoordsX,
		CoordsY:     req.CoordsY,
		SortOrder:   req.SortOrder,
		Status:      1,
	}

	if err := s.spaceRepo.CreateLocation(location); err != nil {
		return nil, err
	}
	return location, nil
}

func (s *SpaceService) GetLocation(id uint) (*domain.SpaceLocation, error) {
	return s.spaceRepo.GetLocationByID(id)
}

func (s *SpaceService) GetAllLocations() ([]domain.SpaceLocation, error) {
	return s.spaceRepo.GetAllLocations()
}

func (s *SpaceService) GetLocationsByType(spaceType domain.SpaceType) ([]domain.SpaceLocation, error) {
	return s.spaceRepo.GetLocationsByType(spaceType)
}

func (s *SpaceService) GetLocationsByFloor(floor int) ([]domain.SpaceLocation, error) {
	return s.spaceRepo.GetLocationsByFloor(floor)
}

func (s *SpaceService) GetLocationTree() ([]domain.SpaceTreeResponse, error) {
	locations, err := s.spaceRepo.GetAllLocations()
	if err != nil {
		return nil, err
	}
	return s.buildLocationTree(nil, locations), nil
}

func (s *SpaceService) buildLocationTree(parentID *uint, locations []domain.SpaceLocation) []domain.SpaceTreeResponse {
	var result []domain.SpaceTreeResponse
	for _, loc := range locations {
		isMatch := false
		if parentID == nil && loc.ParentID == nil {
			isMatch = true
		} else if parentID != nil && loc.ParentID != nil && *parentID == *loc.ParentID {
			isMatch = true
		}

		if isMatch {
			treeNode := domain.SpaceTreeResponse{
				SpaceResponse: *loc.ToResponse(),
				Children:      s.buildLocationTree(&loc.ID, locations),
			}
			result = append(result, treeNode)
		}
	}
	return result
}

func (s *SpaceService) UpdateLocation(id uint, req *domain.CreateSpaceRequest) (*domain.SpaceLocation, error) {
	location, err := s.spaceRepo.GetLocationByID(id)
	if err != nil {
		return nil, err
	}

	if req.Code != location.Code {
		existing, _ := s.spaceRepo.GetLocationByCode(req.Code)
		if existing != nil && existing.ID > 0 && existing.ID != id {
			return nil, errors.New("空间代码已存在")
		}
	}

	location.Name = req.Name
	location.Code = req.Code
	location.Type = domain.SpaceType(req.Type)
	location.ParentID = req.ParentID
	location.GalleryID = req.GalleryID
	location.Floor = req.Floor
	location.Description = req.Description
	location.MapData = req.MapData
	location.CoordsX = req.CoordsX
	location.CoordsY = req.CoordsY
	location.SortOrder = req.SortOrder

	if err := s.spaceRepo.UpdateLocation(location); err != nil {
		return nil, err
	}
	return location, nil
}

func (s *SpaceService) DeleteLocation(id uint) error {
	children, _ := s.spaceRepo.GetLocationsByParent(id)
	if len(children) > 0 {
		return errors.New("该空间下还有子空间，无法删除")
	}
	return s.spaceRepo.DeleteLocation(id)
}

func (s *SpaceService) CreateMapElement(req *domain.CreateMapElementRequest) (*domain.MapElement, error) {
	element := &domain.MapElement{
		LocationID:    req.LocationID,
		ElementType:   req.ElementType,
		ElementID:     req.ElementID,
		ElementSubType: req.ElementSubType,
		Name:          req.Name,
		CoordsX:       req.CoordsX,
		CoordsY:       req.CoordsY,
		Width:         req.Width,
		Height:        req.Height,
		Rotation:      req.Rotation,
		Style:         req.Style,
		Icon:          req.Icon,
		Status:        1,
	}

	if err := s.spaceRepo.CreateMapElement(element); err != nil {
		return nil, err
	}
	return element, nil
}

func (s *SpaceService) GetMapElement(id uint) (*domain.MapElement, error) {
	return s.spaceRepo.GetMapElementByID(id)
}

func (s *SpaceService) GetMapElementsByLocation(locationID uint) ([]domain.MapElement, error) {
	return s.spaceRepo.GetMapElementsByLocation(locationID)
}

func (s *SpaceService) UpdateMapElement(id uint, req *domain.CreateMapElementRequest) (*domain.MapElement, error) {
	element, err := s.spaceRepo.GetMapElementByID(id)
	if err != nil {
		return nil, err
	}

	element.LocationID = req.LocationID
	element.ElementType = req.ElementType
	element.ElementID = req.ElementID
	element.ElementSubType = req.ElementSubType
	element.Name = req.Name
	element.CoordsX = req.CoordsX
	element.CoordsY = req.CoordsY
	element.Width = req.Width
	element.Height = req.Height
	element.Rotation = req.Rotation
	element.Style = req.Style
	element.Icon = req.Icon

	if err := s.spaceRepo.UpdateMapElement(element); err != nil {
		return nil, err
	}
	return element, nil
}

func (s *SpaceService) DeleteMapElement(id uint) error {
	return s.spaceRepo.DeleteMapElement(id)
}

func (s *SpaceService) BatchUpdateMapElements(elements []domain.MapElement) error {
	return s.spaceRepo.BatchUpdateMapElements(elements)
}
