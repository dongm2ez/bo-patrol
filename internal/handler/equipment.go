package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/bo-patrol/internal/domain"
	"github.com/bo-patrol/internal/pkg/db"
	"github.com/bo-patrol/internal/pkg/response"
	"github.com/bo-patrol/internal/repository"
	"github.com/bo-patrol/internal/service"
)

var equipService *service.EquipmentService

func initEquipment() {
	equipmentRepo := repository.NewEquipmentRepository(db.DB)
	equipService = service.NewEquipmentService(equipmentRepo)
}

func CreateEquipment(c *gin.Context) {
	var req domain.CreateEquipmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, err.Error())
		return
	}

	equipment, err := equipService.CreateEquipment(&req)
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, equipment)
}

func GetEquipment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	equipment, err := equipService.GetEquipment(uint(id))
	if err != nil {
		response.Fail(c, http.StatusNotFound, "设备不存在")
		return
	}

	response.Success(c, equipment)
}

func ListEquipments(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))

	filters := make(map[string]interface{})
	if deptId := c.Query("departmentId"); deptId != "" {
		filters["department_id"] = deptId
	}
	if category := c.Query("category"); category != "" {
		filters["category"] = category
	}
	if status := c.Query("status"); status != "" {
		filters["status"] = status
	}

	equipments, total, err := equipService.GetAllEquipments(page, pageSize, filters)
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, gin.H{
		"list":     equipments,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}

func GetEquipmentByDepartment(c *gin.Context) {
	deptId, _ := strconv.Atoi(c.Param("departmentId"))
	equipments, err := equipService.GetEquipmentsByDepartment(uint(deptId))
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, gin.H{
		"list":  equipments,
		"total": len(equipments),
	})
}

func UpdateEquipment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req domain.CreateEquipmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, err.Error())
		return
	}

	equipment, err := equipService.UpdateEquipment(uint(id), &req)
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, equipment)
}

func DeleteEquipment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := equipService.DeleteEquipment(uint(id)); err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, nil)
}

func GetEquipmentCategories(c *gin.Context) {
	categories := []map[string]interface{}{}
	for k, v := range domain.EquipmentCategoryNames {
		categories = append(categories, map[string]interface{}{
			"value": k,
			"label": v,
		})
	}
	response.Success(c, categories)
}

func CreateSupplier(c *gin.Context) {
	var req domain.CreateSupplierRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, err.Error())
		return
	}

	supplier, err := equipService.CreateSupplier(&req)
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, supplier)
}

func ListSuppliers(c *gin.Context) {
	suppliers, err := equipService.GetAllSuppliers()
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, gin.H{
		"list":  suppliers,
		"total": len(suppliers),
	})
}

func GetMaintenanceRecords(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	records, err := equipService.GetMaintenanceRecords(uint(id), 0)
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, gin.H{
		"list":  records,
		"total": len(records),
	})
}
