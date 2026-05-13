package main

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/bo-patrol/internal/domain"
	"github.com/bo-patrol/internal/pkg/config"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	dsn := config.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	log.Println("开始迁移数据库...")

	err = db.AutoMigrate(
		&domain.User{},
		&domain.Department{},
		&domain.Supplier{},
		&domain.SpaceLocation{},
		&domain.MapElement{},
		&domain.Gallery{},
		&domain.Exhibit{},
		&domain.Equipment{},
		&domain.MaintenanceRecord{},
		&domain.PatrolRoute{},
		&domain.RoutePoint{},
		&domain.PatrolTask{},
		&domain.PatrolRecord{},
		&domain.EnvironmentRecord{},
		&domain.ExhibitCheckRecord{},
		&domain.SecurityCheckRecord{},
		&domain.Ticket{},
	)
	if err != nil {
		log.Fatalf("迁移失败: %v", err)
	}

	log.Println("数据库迁移完成!")

	if err := seedData(db); err != nil {
		log.Fatalf("初始化数据失败: %v", err)
	}

	log.Println("数据初始化完成!")
}

func seedData(db *gorm.DB) error {
	var count int64
	if err := db.Model(&domain.User{}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		log.Println("用户数据已存在，跳过初始化")
		return nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	admin := &domain.User{
		Username: "admin",
		Password: string(hashedPassword),
		Name:     "系统管理员",
		Phone:    "13800138000",
		Email:    "admin@example.com",
		Role:     domain.RoleAdmin,
		Status:   1,
	}

	if err := db.Create(admin).Error; err != nil {
		return err
	}

	inspectorPass, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	inspector := &domain.User{
		Username: "inspector",
		Password: string(inspectorPass),
		Name:     "巡检员小王",
		Phone:    "13800138001",
		Email:    "inspector@example.com",
		Role:     domain.RoleInspector,
		Status:   1,
	}

	if err := db.Create(inspector).Error; err != nil {
		return err
	}

	repairPass, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	repair := &domain.User{
		Username: "repair",
		Password: string(repairPass),
		Name:     "维修员小李",
		Phone:    "13800138002",
		Email:    "repair@example.com",
		Role:     domain.RoleRepair,
		Status:   1,
	}

	if err := db.Create(repair).Error; err != nil {
		return err
	}

	log.Println("初始化用户数据完成!")
	log.Println("  - 管理员: admin / admin123")
	log.Println("  - 巡检员: inspector / 123456")
	log.Println("  - 维修员: repair / 123456")

	if err := seedDepartmentData(db); err != nil {
		return err
	}

	if err := seedSpaceData(db); err != nil {
		return err
	}

	if err := seedMuseumData(db); err != nil {
		return err
	}

	return nil
}

func seedDepartmentData(db *gorm.DB) error {
	var count int64
	if err := db.Model(&domain.Department{}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		log.Println("科室数据已存在，跳过初始化")
		return nil
	}

	depts := []domain.Department{
		{
			Name:        "总务处",
			Code:        "DEPT-001",
			Type:        domain.DeptTypeAdmin,
			Description: "博物馆总务管理部门",
			SortOrder:   1,
			Status:      1,
		},
		{
			Name:        "水电科",
			Code:        "DEPT-002",
			Type:        domain.DeptTypeWater,
			Description: "负责给排水和空调系统维护",
			SortOrder:   2,
			Status:      1,
		},
		{
			Name:        "强电科",
			Code:        "DEPT-003",
			Type:        domain.DeptTypeElectric,
			Description: "负责电力系统和仪表维护",
			SortOrder:   3,
			Status:      1,
		},
		{
			Name:        "电梯科",
			Code:        "DEPT-004",
			Type:        domain.DeptTypeElevator,
			Description: "负责电梯和扶梯维护",
			SortOrder:   4,
			Status:      1,
		},
		{
			Name:        "楼宇自动化科",
			Code:        "DEPT-005",
			Type:        domain.DeptTypeBAS,
			Description: "负责监控、传感器系统维护",
			SortOrder:   5,
			Status:      1,
		},
		{
			Name:        "音响设备科",
			Code:        "DEPT-006",
			Type:        domain.DeptTypeAudio,
			Description: "负责音响、话筒设备维护",
			SortOrder:   6,
			Status:      1,
		},
		{
			Name:        "土木工程改造科",
			Code:        "DEPT-007",
			Type:        domain.DeptTypeCivil,
			Description: "负责建筑维护和工程改造",
			SortOrder:   7,
			Status:      1,
		},
		{
			Name:        "财务科",
			Code:        "DEPT-008",
			Type:        domain.DeptTypeFinance,
			Description: "负责财务管理和耗材采购",
			SortOrder:   8,
			Status:      1,
		},
	}

	for i := range depts {
		if err := db.Create(&depts[i]).Error; err != nil {
			return err
		}
	}
	log.Println("初始化科室数据完成!")
	log.Printf("  - 已创建 %d 个科室", len(depts))
	return nil
}

func seedSpaceData(db *gorm.DB) error {
	var count int64
	if err := db.Model(&domain.SpaceLocation{}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		log.Println("空间数据已存在，跳过初始化")
		return nil
	}

	building := domain.SpaceLocation{
		Name:        "博物馆主楼",
		Code:        "BLD-001",
		Type:        domain.SpaceBuilding,
		Floor:       0,
		Description: "博物馆主体建筑",
		Status:      1,
	}
	if err := db.Create(&building).Error; err != nil {
		return err
	}

	floors := []domain.SpaceLocation{
		{
			Name:        "地下室",
			Code:        "FLR-B1",
			Type:        domain.SpaceFloor,
			ParentID:    &building.ID,
			Floor:       -1,
			Description: "地下一层，设备机房区域",
			Status:      1,
		},
		{
			Name:        "一层",
			Code:        "FLR-01",
			Type:        domain.SpaceFloor,
			ParentID:    &building.ID,
			Floor:       1,
			Description: "一楼，主要展厅区域",
			Status:      1,
		},
		{
			Name:        "二层",
			Code:        "FLR-02",
			Type:        domain.SpaceFloor,
			ParentID:    &building.ID,
			Floor:       2,
			Description: "二楼，办公和临时展厅",
			Status:      1,
		},
	}

	for i := range floors {
		if err := db.Create(&floors[i]).Error; err != nil {
			return err
		}
	}

	rooms := []domain.SpaceLocation{
		{
			Name:        "配电室",
			Code:        "ROM-B1-001",
			Type:        domain.SpaceRoom,
			ParentID:    &floors[0].ID,
			Floor:       -1,
			Description: "B1层配电室",
			Status:      1,
		},
		{
			Name:        "空调机房",
			Code:        "ROM-B1-002",
			Type:        domain.SpaceRoom,
			ParentID:    &floors[0].ID,
			Floor:       -1,
			Description: "B1层空调机房",
			Status:      1,
		},
		{
			Name:        "监控中心",
			Code:        "ROM-01-001",
			Type:        domain.SpaceRoom,
			ParentID:    &floors[1].ID,
			Floor:       1,
			Description: "一楼安防监控中心",
			Status:      1,
		},
		{
			Name:        "消防控制室",
			Code:        "ROM-01-002",
			Type:        domain.SpaceRoom,
			ParentID:    &floors[1].ID,
			Floor:       1,
			Description: "一楼消防控制室",
			Status:      1,
		},
	}

	for i := range rooms {
		if err := db.Create(&rooms[i]).Error; err != nil {
			return err
		}
	}

	log.Println("初始化空间数据完成!")
	log.Printf("  - 已创建 %d 个空间节点", 1+len(floors)+len(rooms))
	return nil
}

func seedMuseumData(db *gorm.DB) error {
	var galleryCount int64
	if err := db.Model(&domain.Gallery{}).Count(&galleryCount).Error; err != nil {
		return err
	}
	if galleryCount > 0 {
		log.Println("博物馆数据已存在，跳过初始化")
		return nil
	}

	galleries := []domain.Gallery{
		{
			Name:           "一号展厅 - 古代文物",
			Code:           "GALLERY-001",
			Zone:           domain.GalleryZoneExhibition,
			Floor:          1,
			Description:    "展示珍贵古代文物，包括青铜器、玉器等",
			Area:           200.0,
			TemperatureMin: 18.0,
			TemperatureMax: 22.0,
			HumidityMin:    45.0,
			HumidityMax:    55.0,
		},
		{
			Name:           "二号展厅 - 书画艺术",
			Code:           "GALLERY-002",
			Zone:           domain.GalleryZoneExhibition,
			Floor:          1,
			Description:    "展示古代书画作品，对温湿度要求严格",
			Area:           180.0,
			TemperatureMin: 18.0,
			TemperatureMax: 20.0,
			HumidityMin:    50.0,
			HumidityMax:    55.0,
		},
		{
			Name:           "文物库房 A区",
			Code:           "STORAGE-A01",
			Zone:           domain.GalleryZoneStorage,
			Floor:          0,
			Description:    "一级文物库房，24小时监控",
			Area:           100.0,
			TemperatureMin: 18.0,
			TemperatureMax: 20.0,
			HumidityMin:    45.0,
			HumidityMax:    50.0,
		},
	}

	for i := range galleries {
		if err := db.Create(&galleries[i]).Error; err != nil {
			return err
		}
	}
	log.Println("初始化展厅数据完成!")

	exhibits := []domain.Exhibit{
		{
			Name:         "商代青铜鼎",
			Code:         "EX-0001",
			GalleryID:    galleries[0].ID,
			Level:        domain.ExhibitLevelFirst,
			Category:     "青铜器",
			Material:     "青铜",
			Era:          "商代",
			Status:       domain.ExhibitStatusOnDisplay,
			LocationDesc: "一号展厅中央展柜",
		},
		{
			Name:         "宋代山水画轴",
			Code:         "EX-0002",
			GalleryID:    galleries[1].ID,
			Level:        domain.ExhibitLevelNational,
			Category:     "书画",
			Material:     "绢本",
			Era:          "宋代",
			Status:       domain.ExhibitStatusOnDisplay,
			LocationDesc: "二号展厅北墙展柜",
		},
		{
			Name:         "清代青花瓷瓶",
			Code:         "EX-0003",
			GalleryID:    galleries[0].ID,
			Level:        domain.ExhibitLevelSecond,
			Category:     "瓷器",
			Material:     "瓷",
			Era:          "清代",
			Status:       domain.ExhibitStatusOnDisplay,
			LocationDesc: "一号展厅东侧展柜",
		},
	}

	for i := range exhibits {
		if err := db.Create(&exhibits[i]).Error; err != nil {
			return err
		}
	}
	log.Println("初始化展品数据完成!")
	log.Printf("  - 已创建 %d 个展厅, %d 件展品", len(galleries), len(exhibits))

	return nil
}
