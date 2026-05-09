package domain

import (
	"time"
)

type SpaceType string

const (
	SpaceBuilding SpaceType = "building"
	SpaceFloor    SpaceType = "floor"
	SpaceZone     SpaceType = "zone"
	SpaceRoom     SpaceType = "room"
	SpaceArea     SpaceType = "area"
)

type SpaceLocation struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"size:200;not null"`
	Code        string    `gorm:"size:100;uniqueIndex"`
	Type        SpaceType `gorm:"size:50;default:'room'"`
	ParentID    *uint     `gorm:"index"`
	Parent      *SpaceLocation `gorm:"foreignKey:ParentID"`
	GalleryID   *uint     `gorm:"index"`
	Gallery     *Gallery  `gorm:"foreignKey:GalleryID"`
	Floor       int
	Description string    `gorm:"type:text"`
	MapData     string    `gorm:"type:text"`
	CoordsX     float64
	CoordsY     float64
	SortOrder   int       `gorm:"default:0"`
	Status      int       `gorm:"default:1"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type MapElement struct {
	ID            uint    `gorm:"primaryKey"`
	LocationID    uint    `gorm:"not null;index"`
	Location      SpaceLocation `gorm:"foreignKey:LocationID"`
	ElementType   string  `gorm:"size:50"`
	ElementID     *uint   `gorm:"index"`
	ElementSubType string `gorm:"size:50"`
	Name          string  `gorm:"size:200"`
	CoordsX       float64
	CoordsY       float64
	Width         float64
	Height        float64
	Rotation      float64
	Style         string  `gorm:"type:text"`
	Icon          string  `gorm:"size:200"`
	SortOrder     int     `gorm:"default:0"`
	Status        int     `gorm:"default:1"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type CreateSpaceRequest struct {
	Name        string    `json:"name" binding:"required"`
	Code        string    `json:"code" binding:"required"`
	Type        SpaceType `json:"type"`
	ParentID    *uint     `json:"parentId"`
	GalleryID   *uint     `json:"galleryId"`
	Floor       int       `json:"floor"`
	Description string    `json:"description"`
	MapData     string    `json:"mapData"`
	CoordsX     float64   `json:"coordsX"`
	CoordsY     float64   `json:"coordsY"`
	SortOrder   int       `json:"sortOrder"`
}

type SpaceResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	Type        SpaceType `json:"type"`
	TypeName    string    `json:"typeName"`
	ParentID    *uint     `json:"parentId"`
	ParentName  string    `json:"parentName"`
	GalleryID   *uint     `json:"galleryId"`
	GalleryName string    `json:"galleryName"`
	Floor       int       `json:"floor"`
	Description string    `json:"description"`
	MapData     string    `json:"mapData"`
	CoordsX     float64   `json:"coordsX"`
	CoordsY     float64   `json:"coordsY"`
	SortOrder   int       `json:"sortOrder"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
}

type SpaceTreeResponse struct {
	SpaceResponse
	Children []SpaceTreeResponse `json:"children"`
}

type CreateMapElementRequest struct {
	LocationID    uint    `json:"locationId" binding:"required"`
	ElementType   string  `json:"elementType"`
	ElementID     *uint   `json:"elementId"`
	ElementSubType string `json:"elementSubType"`
	Name          string  `json:"name"`
	CoordsX       float64 `json:"coordsX"`
	CoordsY       float64 `json:"coordsY"`
	Width         float64 `json:"width"`
	Height        float64 `json:"height"`
	Rotation      float64 `json:"rotation"`
	Style         string  `json:"style"`
	Icon          string  `json:"icon"`
}

type MapElementResponse struct {
	ID            uint    `json:"id"`
	LocationID    uint    `json:"locationId"`
	LocationName  string  `json:"locationName"`
	ElementType   string  `json:"elementType"`
	ElementID     *uint   `json:"elementId"`
	ElementSubType string `json:"elementSubType"`
	Name          string  `json:"name"`
	CoordsX       float64 `json:"coordsX"`
	CoordsY       float64 `json:"coordsY"`
	Width         float64 `json:"width"`
	Height        float64 `json:"height"`
	Rotation      float64 `json:"rotation"`
	Style         string  `json:"style"`
	Icon          string  `json:"icon"`
	SortOrder     int     `json:"sortOrder"`
	EquipmentInfo *EquipmentResponse `json:"equipmentInfo,omitempty"`
	PointInfo     *PointResponse `json:"pointInfo,omitempty"`
}

func (s *SpaceLocation) ToResponse() *SpaceResponse {
	resp := &SpaceResponse{
		ID:          s.ID,
		Name:        s.Name,
		Code:        s.Code,
		Type:        s.Type,
		TypeName:    SpaceTypeNames[s.Type],
		ParentID:    s.ParentID,
		GalleryID:   s.GalleryID,
		Floor:       s.Floor,
		Description: s.Description,
		MapData:     s.MapData,
		CoordsX:     s.CoordsX,
		CoordsY:     s.CoordsY,
		SortOrder:   s.SortOrder,
		Status:      s.Status,
		CreatedAt:   s.CreatedAt,
	}

	if s.Parent != nil {
		resp.ParentName = s.Parent.Name
	}

	if s.Gallery != nil {
		resp.GalleryName = s.Gallery.Name
	}

	return resp
}

func (m *MapElement) ToResponse() *MapElementResponse {
	return &MapElementResponse{
		ID:             m.ID,
		LocationID:     m.LocationID,
		LocationName:   m.Location.Name,
		ElementType:    m.ElementType,
		ElementID:      m.ElementID,
		ElementSubType: m.ElementSubType,
		Name:           m.Name,
		CoordsX:        m.CoordsX,
		CoordsY:        m.CoordsY,
		Width:          m.Width,
		Height:         m.Height,
		Rotation:       m.Rotation,
		Style:          m.Style,
		Icon:           m.Icon,
		SortOrder:      m.SortOrder,
	}
}

var SpaceTypeNames = map[SpaceType]string{
	SpaceBuilding: "建筑",
	SpaceFloor:    "楼层",
	SpaceZone:     "区域",
	SpaceRoom:     "房间",
	SpaceArea:     "区域",
}

var ElementTypes = []string{
	"equipment",
	"patrol_point",
	"exhibit",
	"door",
	"window",
	"fire_extinguisher",
	"camera",
	"sensor",
	"note",
}

var ElementTypeNames = map[string]string{
	"equipment":       "设备",
	"patrol_point":    "巡检点",
	"exhibit":         "展品",
	"door":            "门",
	"window":          "窗",
	"fire_extinguisher": "灭火器",
	"camera":          "摄像头",
	"sensor":          "传感器",
	"note":            "标注",
}
