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

var spaceService *service.SpaceService

func initSpace() {
	spaceRepo := repository.NewSpaceRepository(db.DB)
	spaceService = service.NewSpaceService(spaceRepo)
}

func CreateLocation(c *gin.Context) {
	var req domain.CreateSpaceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, err.Error())
		return
	}

	location, err := spaceService.CreateLocation(&req)
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, location)
}

func GetLocation(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	location, err := spaceService.GetLocation(uint(id))
	if err != nil {
		response.Fail(c, http.StatusNotFound, "空间不存在")
		return
	}

	response.Success(c, location)
}

func ListLocations(c *gin.Context) {
	spaceType := c.Query("type")
	floor, _ := strconv.Atoi(c.Query("floor"))

	var locations []domain.SpaceLocation
	var err error

	if spaceType != "" {
		locations, err = spaceService.GetLocationsByType(domain.SpaceType(spaceType))
	} else if floor != 0 {
		locations, err = spaceService.GetLocationsByFloor(floor)
	} else {
		locations, err = spaceService.GetAllLocations()
	}

	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, gin.H{
		"list":  locations,
		"total": len(locations),
	})
}

func GetLocationTree(c *gin.Context) {
	tree := spaceService.GetLocationTree()
	response.Success(c, tree)
}

func UpdateLocation(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req domain.CreateSpaceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, err.Error())
		return
	}

	location, err := spaceService.UpdateLocation(uint(id), &req)
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, location)
}

func DeleteLocation(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := spaceService.DeleteLocation(uint(id)); err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, nil)
}

func GetSpaceTypes(c *gin.Context) {
	types := []map[string]interface{}{
		{"value": domain.SpaceBuilding, "label": "建筑"},
		{"value": domain.SpaceFloor, "label": "楼层"},
		{"value": domain.SpaceZone, "label": "区域"},
		{"value": domain.SpaceRoom, "label": "房间"},
	}
	response.Success(c, types)
}

func CreateMapElement(c *gin.Context) {
	var req domain.CreateMapElementRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, err.Error())
		return
	}

	element, err := spaceService.CreateMapElement(&req)
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, element)
}

func GetMapElement(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	element, err := spaceService.GetMapElement(uint(id))
	if err != nil {
		response.Fail(c, http.StatusNotFound, "地图元素不存在")
		return
	}

	response.Success(c, element)
}

func GetMapElementsByLocation(c *gin.Context) {
	locationId, _ := strconv.Atoi(c.Param("locationId"))
	elements, err := spaceService.GetMapElementsByLocation(uint(locationId))
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, gin.H{
		"list":  elements,
		"total": len(elements),
	})
}

func UpdateMapElement(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req domain.CreateMapElementRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, err.Error())
		return
	}

	element, err := spaceService.UpdateMapElement(uint(id), &req)
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, element)
}

func DeleteMapElement(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := spaceService.DeleteMapElement(uint(id)); err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, nil)
}

func BatchUpdateMapElements(c *gin.Context) {
	var elements []domain.MapElement
	if err := c.ShouldBindJSON(&elements); err != nil {
		response.ParamError(c, err.Error())
		return
	}

	if err := spaceService.BatchUpdateMapElements(elements); err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, nil)
}

func GetElementTypes(c *gin.Context) {
	types := []map[string]interface{}{}
	for _, t := range domain.ElementTypes {
		types = append(types, map[string]interface{}{
			"value": t,
			"label": domain.ElementTypeNames[t],
		})
	}
	response.Success(c, types)
}
