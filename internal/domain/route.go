package domain

import (
	"time"
)

type PatrolRoute struct {
	ID          uint        `gorm:"primaryKey"`
	Name        string      `gorm:"size:100;not null"`
	Description string      `gorm:"type:text"`
	Zone        string      `gorm:"size:50"`
	Status      int         `gorm:"default:1"`
	CreatedBy   uint        `gorm:"not null"`
	Points      []RoutePoint `gorm:"foreignKey:RouteID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type PointType string

const (
	PointTypeGallery   PointType = "gallery"
	PointTypeExhibit   PointType = "exhibit"
	PointTypePublic    PointType = "public"
	PointTypeEquipment PointType = "equipment"
	PointTypeSecurity  PointType = "security"
)

type RoutePoint struct {
	ID           uint      `gorm:"primaryKey"`
	RouteID      uint      `gorm:"not null;index"`
	Name         string    `gorm:"size:100;not null"`
	PointType    PointType `gorm:"size:50;default:'public'"`
	GalleryID    *uint     `gorm:"index"`
	Gallery      Gallery   `gorm:"foreignKey:GalleryID"`
	ExhibitID    *uint     `gorm:"index"`
	Exhibit      Exhibit   `gorm:"foreignKey:ExhibitID"`
	QRCode       string    `gorm:"size:100;uniqueIndex"`
	Latitude     float64
	Longitude    float64
	CheckItems   string    `gorm:"type:text"`
	CheckEnv     bool      `gorm:"default:false"`
	CheckExhibit bool      `gorm:"default:false"`
	CheckSecurity bool     `gorm:"default:false"`
	Status       int       `gorm:"default:1"`
	Sort         int       `gorm:"default:0"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type CreateRouteRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Zone        string `json:"zone"`
}

type CreatePointRequest struct {
	Name          string    `json:"name" binding:"required"`
	PointType     PointType `json:"pointType"`
	GalleryID     *uint     `json:"galleryId"`
	ExhibitID     *uint     `json:"exhibitId"`
	Latitude      float64   `json:"latitude"`
	Longitude     float64   `json:"longitude"`
	CheckItems    string    `json:"checkItems"`
	CheckEnv      bool      `json:"checkEnv"`
	CheckExhibit  bool      `json:"checkExhibit"`
	CheckSecurity bool      `json:"checkSecurity"`
	Sort          int       `json:"sort"`
}

type RouteResponse struct {
	ID          uint            `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Zone        string          `json:"zone"`
	Status      int             `json:"status"`
	Points      []PointResponse `json:"points"`
	CreatedAt   time.Time       `json:"createdAt"`
}

type PointResponse struct {
	ID            uint      `json:"id"`
	RouteID       uint      `json:"routeId"`
	Name          string    `json:"name"`
	PointType     PointType `json:"pointType"`
	GalleryID     *uint     `json:"galleryId"`
	GalleryName   string    `json:"galleryName"`
	ExhibitID     *uint     `json:"exhibitId"`
	ExhibitName   string    `json:"exhibitName"`
	QRCode        string    `json:"qrCode"`
	Latitude      float64   `json:"latitude"`
	Longitude     float64   `json:"longitude"`
	CheckItems    string    `json:"checkItems"`
	CheckEnv      bool      `json:"checkEnv"`
	CheckExhibit  bool      `json:"checkExhibit"`
	CheckSecurity bool      `json:"checkSecurity"`
	Sort          int       `json:"sort"`
}

func (r *PatrolRoute) ToResponse() *RouteResponse {
	points := make([]PointResponse, len(r.Points))
	for i, p := range r.Points {
		points[i] = *p.ToResponse()
	}
	return &RouteResponse{
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
		Zone:        r.Zone,
		Status:      r.Status,
		Points:      points,
		CreatedAt:   r.CreatedAt,
	}
}

func (p *RoutePoint) ToResponse() *PointResponse {
	resp := &PointResponse{
		ID:            p.ID,
		RouteID:       p.RouteID,
		Name:          p.Name,
		PointType:     p.PointType,
		GalleryID:     p.GalleryID,
		ExhibitID:     p.ExhibitID,
		QRCode:        p.QRCode,
		Latitude:      p.Latitude,
		Longitude:     p.Longitude,
		CheckItems:    p.CheckItems,
		CheckEnv:      p.CheckEnv,
		CheckExhibit:  p.CheckExhibit,
		CheckSecurity: p.CheckSecurity,
		Sort:          p.Sort,
	}
	if p.Gallery.Name != "" {
		resp.GalleryName = p.Gallery.Name
	}
	if p.Exhibit.Name != "" {
		resp.ExhibitName = p.Exhibit.Name
	}
	return resp
}
