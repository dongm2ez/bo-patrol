package domain

import (
	"time"
)

type EquipmentStatus string

const (
	EquipStatusNormal    EquipmentStatus = "normal"
	EquipStatusFault     EquipmentStatus = "fault"
	EquipStatusMaintenance EquipmentStatus = "maintenance"
	EquipStatusScrapped  EquipmentStatus = "scrapped"
)

type EquipmentCategory string

const (
	EquipCatWaterPump    EquipmentCategory = "water_pump"
	EquipCatWaterPipe    EquipmentCategory = "water_pipe"
	EquipCatWaterTank    EquipmentCategory = "water_tank"
	EquipCatAC           EquipmentCategory = "air_conditioner"
	EquipCatSwitchboard  EquipmentCategory = "switchboard"
	EquipCatMeter        EquipmentCategory = "meter"
	EquipCatElevator     EquipmentCategory = "elevator"
	EquipCatEscalator    EquipmentCategory = "escalator"
	EquipCatCamera       EquipmentCategory = "camera"
	EquipCatSensor       EquipmentCategory = "sensor"
	EquipCatMicrophone   EquipmentCategory = "microphone"
	EquipCatAudioSystem  EquipmentCategory = "audio_system"
	EquipCatOther        EquipmentCategory = "other"
)

type Equipment struct {
	ID             uint             `gorm:"primaryKey"`
	Name           string           `gorm:"size:200;not null"`
	Code           string           `gorm:"size:100;uniqueIndex"`
	Category       EquipmentCategory `gorm:"size:50;default:'other'"`
	Brand          string           `gorm:"size:100"`
	Model          string           `gorm:"size:100"`
	SerialNumber   string           `gorm:"size:100"`
	DepartmentID   uint             `gorm:"not null;index"`
	Department     Department       `gorm:"foreignKey:DepartmentID"`
	LocationID     *uint            `gorm:"index"`
	Location       *SpaceLocation   `gorm:"foreignKey:LocationID"`
	LocationDesc   string           `gorm:"size:500"`
	PurchaseDate   *time.Time
	PurchasePrice  float64
	SupplierID     *uint
	Supplier       *Supplier        `gorm:"foreignKey:SupplierID"`
	MaintenanceVendor string        `gorm:"size:200"`
	MaintenanceStart *time.Time
	MaintenanceEnd   *time.Time
	LastMaintenance  *time.Time
	NextMaintenance  *time.Time
	Status         EquipmentStatus  `gorm:"size:20;default:'normal'"`
	Specs          string           `gorm:"type:text"`
	Remark         string           `gorm:"type:text"`
	QRCode         string           `gorm:"size:100;uniqueIndex"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Supplier struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:200;not null"`
	Code      string `gorm:"size:50;uniqueIndex"`
	Contact   string `gorm:"size:100"`
	Phone     string `gorm:"size:50"`
	Email     string `gorm:"size:100"`
	Address   string `gorm:"size:500"`
	Type      string `gorm:"size:50"`
	Status    int    `gorm:"default:1"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateEquipmentRequest struct {
	Name              string            `json:"name" binding:"required"`
	Code              string            `json:"code" binding:"required"`
	Category          EquipmentCategory `json:"category"`
	Brand             string            `json:"brand"`
	Model             string            `json:"model"`
	SerialNumber      string            `json:"serialNumber"`
	DepartmentID      uint              `json:"departmentId" binding:"required"`
	LocationID        *uint             `json:"locationId"`
	LocationDesc      string            `json:"locationDesc"`
	PurchaseDate      *string           `json:"purchaseDate"`
	PurchasePrice     float64           `json:"purchasePrice"`
	SupplierID        *uint             `json:"supplierId"`
	MaintenanceVendor string            `json:"maintenanceVendor"`
	MaintenanceStart  *string           `json:"maintenanceStart"`
	MaintenanceEnd    *string           `json:"maintenanceEnd"`
	LastMaintenance   *string           `json:"lastMaintenance"`
	NextMaintenance   *string           `json:"nextMaintenance"`
	Status            EquipmentStatus   `json:"status"`
	Specs             string            `json:"specs"`
	Remark            string            `json:"remark"`
}

type EquipmentResponse struct {
	ID                uint              `json:"id"`
	Name              string            `json:"name"`
	Code              string            `json:"code"`
	Category          EquipmentCategory `json:"category"`
	CategoryName      string            `json:"categoryName"`
	Brand             string            `json:"brand"`
	Model             string            `json:"model"`
	SerialNumber      string            `json:"serialNumber"`
	DepartmentID      uint              `json:"departmentId"`
	DepartmentName    string            `json:"departmentName"`
	LocationID        *uint             `json:"locationId"`
	LocationName      string            `json:"locationName"`
	LocationDesc      string            `json:"locationDesc"`
	PurchaseDate      *time.Time        `json:"purchaseDate"`
	PurchasePrice     float64           `json:"purchasePrice"`
	SupplierID        *uint             `json:"supplierId"`
	SupplierName      string            `json:"supplierName"`
	MaintenanceVendor string            `json:"maintenanceVendor"`
	MaintenanceStart  *time.Time        `json:"maintenanceStart"`
	MaintenanceEnd    *time.Time        `json:"maintenanceEnd"`
	LastMaintenance   *time.Time        `json:"lastMaintenance"`
	NextMaintenance   *time.Time        `json:"nextMaintenance"`
	Status            EquipmentStatus   `json:"status"`
	StatusName        string            `json:"statusName"`
	Specs             string            `json:"specs"`
	Remark            string            `json:"remark"`
	QRCode            string            `json:"qrCode"`
	CreatedAt         time.Time         `json:"createdAt"`
}

type CreateSupplierRequest struct {
	Name    string `json:"name" binding:"required"`
	Code    string `json:"code" binding:"required"`
	Contact string `json:"contact"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Type    string `json:"type"`
}

type SupplierResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	Contact   string `json:"contact"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Type      string `json:"type"`
	Status    int    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

type MaintenanceRecord struct {
	ID          uint      `gorm:"primaryKey"`
	EquipmentID uint      `gorm:"not null;index"`
	Equipment   Equipment `gorm:"foreignKey:EquipmentID"`
	Type        string    `gorm:"size:50"`
	Date        time.Time
	Content     string    `gorm:"type:text"`
	Cost        float64
	Operator    string    `gorm:"size:100"`
	Remark      string    `gorm:"type:text"`
	CreatedAt   time.Time
}

func (e *Equipment) ToResponse() *EquipmentResponse {
	resp := &EquipmentResponse{
		ID:                e.ID,
		Name:              e.Name,
		Code:              e.Code,
		Category:          e.Category,
		CategoryName:      EquipmentCategoryNames[e.Category],
		Brand:             e.Brand,
		Model:             e.Model,
		SerialNumber:      e.SerialNumber,
		DepartmentID:      e.DepartmentID,
		LocationID:        e.LocationID,
		LocationDesc:      e.LocationDesc,
		PurchaseDate:      e.PurchaseDate,
		PurchasePrice:     e.PurchasePrice,
		SupplierID:        e.SupplierID,
		MaintenanceVendor: e.MaintenanceVendor,
		MaintenanceStart:  e.MaintenanceStart,
		MaintenanceEnd:    e.MaintenanceEnd,
		LastMaintenance:   e.LastMaintenance,
		NextMaintenance:   e.NextMaintenance,
		Status:            e.Status,
		StatusName:        EquipmentStatusNames[e.Status],
		Specs:             e.Specs,
		Remark:            e.Remark,
		QRCode:            e.QRCode,
		CreatedAt:         e.CreatedAt,
	}

	if e.Department.Name != "" {
		resp.DepartmentName = e.Department.Name
	}

	if e.Location != nil {
		resp.LocationName = e.Location.Name
	}

	if e.Supplier != nil {
		resp.SupplierName = e.Supplier.Name
	}

	return resp
}

func (s *Supplier) ToResponse() *SupplierResponse {
	return &SupplierResponse{
		ID:        s.ID,
		Name:      s.Name,
		Code:      s.Code,
		Contact:   s.Contact,
		Phone:     s.Phone,
		Email:     s.Email,
		Address:   s.Address,
		Type:      s.Type,
		Status:    s.Status,
		CreatedAt: s.CreatedAt,
	}
}

var EquipmentStatusNames = map[EquipmentStatus]string{
	EquipStatusNormal:    "正常",
	EquipStatusFault:     "故障",
	EquipStatusMaintenance: "维保中",
	EquipStatusScrapped:  "已报废",
}

var EquipmentCategoryNames = map[EquipmentCategory]string{
	EquipCatWaterPump:   "水泵",
	EquipCatWaterPipe:   "水管",
	EquipCatWaterTank:   "水箱",
	EquipCatAC:          "空调",
	EquipCatSwitchboard: "配电柜",
	EquipCatMeter:       "电表",
	EquipCatElevator:    "电梯",
	EquipCatEscalator:   "扶梯",
	EquipCatCamera:      "摄像头",
	EquipCatSensor:      "传感器",
	EquipCatMicrophone:  "话筒",
	EquipCatAudioSystem: "音响系统",
	EquipCatOther:       "其他",
}

var DepartmentEquipmentCategories = map[DepartmentType][]EquipmentCategory{
	DeptTypeWater: {
		EquipCatWaterPump,
		EquipCatWaterPipe,
		EquipCatWaterTank,
		EquipCatAC,
	},
	DeptTypeElectric: {
		EquipCatSwitchboard,
		EquipCatMeter,
	},
	DeptTypeElevator: {
		EquipCatElevator,
		EquipCatEscalator,
	},
	DeptTypeBAS: {
		EquipCatCamera,
		EquipCatSensor,
	},
	DeptTypeAudio: {
		EquipCatMicrophone,
		EquipCatAudioSystem,
	},
	DeptTypeCivil: {},
	DeptTypeFinance: {},
	DeptTypeAdmin: {},
	DeptTypeOther: {},
}
