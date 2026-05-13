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

var deptService *service.DepartmentService

func initDepartment() {
	deptRepo := repository.NewDepartmentRepository(db.DB)
	deptService = service.NewDepartmentService(deptRepo)
}

func CreateDepartment(c *gin.Context) {
	var req domain.CreateDepartmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, err.Error())
		return
	}

	dept, err := deptService.CreateDepartment(&req)
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, dept)
}

func GetDepartment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.ParamError(c, "无效的ID参数")
		return
	}
	dept, err := deptService.GetDepartment(uint(id))
	if err != nil {
		response.Fail(c, http.StatusNotFound, "科室不存在")
		return
	}

	response.Success(c, dept)
}

func ListDepartments(c *gin.Context) {
	deptType := c.Query("type")
	var depts []domain.Department
	var err error

	if deptType != "" {
		depts, err = deptService.GetDepartmentsByType(domain.DepartmentType(deptType))
	} else {
		depts, err = deptService.GetAllDepartments()
	}

	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, gin.H{
		"list":  depts,
		"total": len(depts),
	})
}

func GetDepartmentTree(c *gin.Context) {
	tree, err := deptService.GetDepartmentTree()
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, tree)
}

func UpdateDepartment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.ParamError(c, "无效的ID参数")
		return
	}
	var req domain.CreateDepartmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, err.Error())
		return
	}

	dept, err := deptService.UpdateDepartment(uint(id), &req)
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, dept)
}

func DeleteDepartment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.ParamError(c, "无效的ID参数")
		return
	}
	if err := deptService.DeleteDepartment(uint(id)); err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, nil)
}

func GetDepartmentTypes(c *gin.Context) {
	types := []map[string]interface{}{}
	for _, t := range domain.DefaultDepartmentTypes {
		types = append(types, map[string]interface{}{
			"value": t,
			"label": domain.DepartmentTypeNames[t],
		})
	}
	response.Success(c, types)
}
