package domain

import (
	"time"
)

type GalleryZone string

const (
	GalleryZoneExhibition    GalleryZone = "exhibition"
	GalleryZoneStorage       GalleryZone = "storage"
	GalleryZonePublic        GalleryZone = "public"
	GalleryZoneOffice        GalleryZone = "office"
	GalleryZoneEquipment     GalleryZone = "equipment"
)

type ExhibitStatus string

const (
	ExhibitStatusOnDisplay  ExhibitStatus = "on_display"
	ExhibitStatusInStorage  ExhibitStatus = "in_storage"
	ExhibitStatusOnLoan     ExhibitStatus = "on_loan"
	ExhibitStatusInRepair   ExhibitStatus = "in_repair"
)

type ExhibitLevel string

const (
	ExhibitLevelNational    ExhibitLevel = "national"
	ExhibitLevelFirst       ExhibitLevel = "first"
	ExhibitLevelSecond      ExhibitLevel = "second"
	ExhibitLevelThird       ExhibitLevel = "third"
	ExhibitLevelGeneral     ExhibitLevel = "general"
)

type Gallery struct {
	ID          uint        `gorm:"primaryKey"`
	Name        string      `gorm:"size:100;not null"`
	Code        string      `gorm:"size:50;uniqueIndex"`
	Zone        GalleryZone `gorm:"size:50"`
	Floor       int
	Description string      `gorm:"type:text"`
	Area        float64
	TemperatureMin float64
	TemperatureMax float64
	HumidityMin    float64
	HumidityMax    float64
	Status      int         `gorm:"default:1"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Exhibit struct {
	ID          uint          `gorm:"primaryKey"`
	Name        string        `gorm:"size:200;not null"`
	Code        string        `gorm:"size:100;uniqueIndex"`
	GalleryID   uint          `gorm:"index"`
	Gallery     Gallery       `gorm:"foreignKey:GalleryID"`
	Level       ExhibitLevel  `gorm:"size:50"`
	Category    string        `gorm:"size:100"`
	Description string        `gorm:"type:text"`
	Material    string        `gorm:"size:200"`
	Era         string        `gorm:"size:100"`
	Status      ExhibitStatus `gorm:"size:50"`
	LocationDesc string       `gorm:"size:200"`
	ImageURL    string        `gorm:"size:500"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type EnvironmentRecord struct {
	ID          uint      `gorm:"primaryKey"`
	GalleryID   uint      `gorm:"not null;index"`
	Gallery     Gallery   `gorm:"foreignKey:GalleryID"`
	PointID     uint      `gorm:"index"`
	RecordID    uint      `gorm:"index"`
	Temperature float64
	Humidity    float64
	Lux         float64
	CO2         float64
	CheckedBy   uint
	CheckedAt   time.Time
	Remark      string    `gorm:"type:text"`
	CreatedAt   time.Time
}

type ExhibitCheckRecord struct {
	ID          uint      `gorm:"primaryKey"`
	ExhibitID   uint      `gorm:"not null;index"`
	Exhibit     Exhibit   `gorm:"foreignKey:ExhibitID"`
	RecordID    uint      `gorm:"not null;index"`
	IsIntact    bool
	IsInPosition bool
	HasDamage   bool
	DamageType  string    `gorm:"size:100"`
	CaseSealed  bool
	Remark      string    `gorm:"type:text"`
	ImageURLs   string    `gorm:"type:text"`
	CheckedBy   uint
	CheckedAt   time.Time
	CreatedAt   time.Time
}

type SecurityCheckRecord struct {
	ID          uint      `gorm:"primaryKey"`
	PointID     uint      `gorm:"not null;index"`
	RecordID    uint      `gorm:"not null;index"`
	CameraWorking bool
	AlarmWorking  bool
	DoorLocked    bool
	WindowLocked  bool
	FireEquipment bool
	EmergencyLight bool
	Remark        string  `gorm:"type:text"`
	CheckedBy     uint
	CheckedAt     time.Time
	CreatedAt     time.Time
}

type CreateGalleryRequest struct {
	Name        string      `json:"name" binding:"required"`
	Code        string      `json:"code" binding:"required"`
	Zone        GalleryZone `json:"zone"`
	Floor       int         `json:"floor"`
	Description string      `json:"description"`
	Area        float64     `json:"area"`
	TemperatureMin float64  `json:"temperatureMin"`
	TemperatureMax float64  `json:"temperatureMax"`
	HumidityMin    float64  `json:"humidityMin"`
	HumidityMax    float64  `json:"humidityMax"`
}

type CreateExhibitRequest struct {
	Name        string       `json:"name" binding:"required"`
	Code        string       `json:"code" binding:"required"`
	GalleryID   uint         `json:"galleryId"`
	Level       ExhibitLevel `json:"level"`
	Category    string       `json:"category"`
	Description string       `json:"description"`
	Material    string       `json:"material"`
	Era         string       `json:"era"`
	LocationDesc string      `json:"locationDesc"`
}

type MuseumCheckItem struct {
	CheckType    string `json:"checkType"`
	Name         string `json:"name"`
	Standard     string `json:"standard"`
	IsRequired   bool   `json:"isRequired"`
}

var GalleryCheckItems = []MuseumCheckItem{
	{CheckType: "environment", Name: "温度", Standard: "符合展品保存要求", IsRequired: true},
	{CheckType: "environment", Name: "湿度", Standard: "符合展品保存要求", IsRequired: true},
	{CheckType: "environment", Name: "光照度", Standard: "对书画、纺织品等光敏展品无损害", IsRequired: false},
	{CheckType: "security", Name: "监控摄像头", Standard: "运行正常，覆盖完整", IsRequired: true},
	{CheckType: "security", Name: "防盗报警", Standard: "系统正常", IsRequired: true},
	{CheckType: "security", Name: "消防设备", Standard: "完好有效", IsRequired: true},
	{CheckType: "exhibit", Name: "展品完整性", Standard: "展品完好无损坏", IsRequired: true},
	{CheckType: "exhibit", Name: "展品位置", Standard: "位置正确无移位", IsRequired: true},
	{CheckType: "exhibit", Name: "展柜密封", Standard: "展柜密封良好", IsRequired: true},
}

type GalleryResponse struct {
	ID          uint        `json:"id"`
	Name        string      `json:"name"`
	Code        string      `json:"code"`
	Zone        GalleryZone `json:"zone"`
	Floor       int         `json:"floor"`
	Description string      `json:"description"`
	Area        float64     `json:"area"`
	TemperatureMin float64  `json:"temperatureMin"`
	TemperatureMax float64  `json:"temperatureMax"`
	HumidityMin    float64  `json:"humidityMin"`
	HumidityMax    float64  `json:"humidityMax"`
	Status      int         `json:"status"`
	CreatedAt   time.Time   `json:"createdAt"`
}

type ExhibitResponse struct {
	ID          uint          `json:"id"`
	Name        string        `json:"name"`
	Code        string        `json:"code"`
	GalleryID   uint          `json:"galleryId"`
	GalleryName string        `json:"galleryName"`
	Level       ExhibitLevel  `json:"level"`
	Category    string        `json:"category"`
	Description string        `json:"description"`
	Material    string        `json:"material"`
	Era         string        `json:"era"`
	Status      ExhibitStatus `json:"status"`
	LocationDesc string       `json:"locationDesc"`
	ImageURL    string        `json:"imageUrl"`
	CreatedAt   time.Time     `json:"createdAt"`
}

type EnvironmentRecordResponse struct {
	ID          uint      `json:"id"`
	GalleryID   uint      `json:"galleryId"`
	GalleryName string    `json:"galleryName"`
	Temperature float64   `json:"temperature"`
	Humidity    float64   `json:"humidity"`
	Lux         float64   `json:"lux"`
	CO2         float64   `json:"co2"`
	CheckedBy   uint      `json:"checkedBy"`
	CheckedAt   time.Time `json:"checkedAt"`
	Remark      string    `json:"remark"`
}

func (g *Gallery) ToResponse() *GalleryResponse {
	return &GalleryResponse{
		ID:              g.ID,
		Name:            g.Name,
		Code:            g.Code,
		Zone:            g.Zone,
		Floor:           g.Floor,
		Description:     g.Description,
		Area:            g.Area,
		TemperatureMin:  g.TemperatureMin,
		TemperatureMax:  g.TemperatureMax,
		HumidityMin:     g.HumidityMin,
		HumidityMax:     g.HumidityMax,
		Status:          g.Status,
		CreatedAt:       g.CreatedAt,
	}
}

func (e *Exhibit) ToResponse() *ExhibitResponse {
	return &ExhibitResponse{
		ID:           e.ID,
		Name:         e.Name,
		Code:         e.Code,
		GalleryID:    e.GalleryID,
		GalleryName:  e.Gallery.Name,
		Level:        e.Level,
		Category:     e.Category,
		Description:  e.Description,
		Material:     e.Material,
		Era:          e.Era,
		Status:       e.Status,
		LocationDesc: e.LocationDesc,
		ImageURL:     e.ImageURL,
		CreatedAt:    e.CreatedAt,
	}
}

func (e *EnvironmentRecord) ToResponse() *EnvironmentRecordResponse {
	return &EnvironmentRecordResponse{
		ID:          e.ID,
		GalleryID:   e.GalleryID,
		GalleryName: e.Gallery.Name,
		Temperature: e.Temperature,
		Humidity:    e.Humidity,
		Lux:         e.Lux,
		CO2:         e.CO2,
		CheckedBy:   e.CheckedBy,
		CheckedAt:   e.CheckedAt,
		Remark:      e.Remark,
	}
}
