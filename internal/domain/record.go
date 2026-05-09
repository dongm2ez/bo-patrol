package domain

import (
	"time"
)

type RecordStatus string
type CheckMethod string

const (
	RecordStatusNormal  RecordStatus = "normal"
	RecordStatusAbnormal RecordStatus = "abnormal"
	RecordStatusMissed  RecordStatus = "missed"
)

const (
	CheckMethodQRCode CheckMethod = "qr_code"
	CheckMethodNFC    CheckMethod = "nfc"
	CheckMethodGPS    CheckMethod = "gps"
)

type PatrolRecord struct {
	ID          uint         `gorm:"primaryKey"`
	TaskID      uint         `gorm:"not null;index"`
	PointID     uint         `gorm:"not null"`
	InspectorID uint         `gorm:"not null;index"`
	Status      RecordStatus `gorm:"size:20;default:'normal'"`
	CheckTime   time.Time    `gorm:"not null"`
	CheckMethod CheckMethod  `gorm:"size:20;default:'qr_code'"`
	CheckResult string       `gorm:"type:text"`
	Images      string       `gorm:"type:text"`
	AbnormalDesc string      `gorm:"type:text"`
	TicketID    *uint
	Task        PatrolTask   `gorm:"foreignKey:TaskID"`
	Point       RoutePoint   `gorm:"foreignKey:PointID"`
	Inspector   User         `gorm:"foreignKey:InspectorID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type EnvironmentData struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
	Lux         float64 `json:"lux"`
	CO2         float64 `json:"co2"`
	Remark      string  `json:"remark"`
}

type ExhibitCheckData struct {
	ExhibitID   uint   `json:"exhibitId" binding:"required"`
	IsIntact    bool   `json:"isIntact"`
	IsInPosition bool  `json:"isInPosition"`
	HasDamage   bool   `json:"hasDamage"`
	DamageType  string `json:"damageType"`
	CaseSealed  bool   `json:"caseSealed"`
	Remark      string `json:"remark"`
	ImageURLs   string `json:"imageUrls"`
}

type SecurityCheckData struct {
	CameraWorking  bool `json:"cameraWorking"`
	AlarmWorking   bool `json:"alarmWorking"`
	DoorLocked     bool `json:"doorLocked"`
	WindowLocked   bool `json:"windowLocked"`
	FireEquipment  bool `json:"fireEquipment"`
	EmergencyLight bool `json:"emergencyLight"`
	Remark         string `json:"remark"`
}

type CreateRecordRequest struct {
	TaskID        uint              `json:"taskId" binding:"required"`
	PointID       uint              `json:"pointId" binding:"required"`
	QRCode        string            `json:"qrCode" binding:"required"`
	CheckResult   string            `json:"checkResult"`
	Images        string            `json:"images"`
	IsAbnormal    bool              `json:"isAbnormal"`
	AbnormalDesc  string            `json:"abnormalDesc"`
	Environment   *EnvironmentData  `json:"environment"`
	ExhibitChecks []ExhibitCheckData `json:"exhibitChecks"`
	Security      *SecurityCheckData `json:"security"`
}

type RecordResponse struct {
	ID             uint              `json:"id"`
	TaskID         uint              `json:"taskId"`
	TaskName       string            `json:"taskName"`
	PointID        uint              `json:"pointId"`
	PointName      string            `json:"pointName"`
	PointType      PointType         `json:"pointType"`
	InspectorID    uint              `json:"inspectorId"`
	InspectorName  string            `json:"inspectorName"`
	Status         RecordStatus      `json:"status"`
	CheckTime      time.Time         `json:"checkTime"`
	CheckMethod    CheckMethod       `json:"checkMethod"`
	CheckResult    string            `json:"checkResult"`
	Images         string            `json:"images"`
	AbnormalDesc   string            `json:"abnormalDesc"`
	TicketID       *uint             `json:"ticketId"`
	Environment    *EnvironmentData  `json:"environment"`
	ExhibitChecks  []ExhibitCheckData `json:"exhibitChecks"`
	Security       *SecurityCheckData `json:"security"`
}

func (r *PatrolRecord) ToResponse() *RecordResponse {
	resp := &RecordResponse{
		ID:           r.ID,
		TaskID:       r.TaskID,
		PointID:      r.PointID,
		InspectorID:  r.InspectorID,
		Status:       r.Status,
		CheckTime:    r.CheckTime,
		CheckMethod:  r.CheckMethod,
		CheckResult:  r.CheckResult,
		Images:       r.Images,
		AbnormalDesc: r.AbnormalDesc,
		TicketID:     r.TicketID,
	}
	if r.Task.Name != "" {
		resp.TaskName = r.Task.Name
	}
	if r.Point.Name != "" {
		resp.PointName = r.Point.Name
		resp.PointType = r.Point.PointType
	}
	if r.Inspector.Name != "" {
		resp.InspectorName = r.Inspector.Name
	}
	return resp
}
