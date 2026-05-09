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

var museumService *service.MuseumService

func initMuseum() {
	museumRepo := repository.NewMuseumRepository(db.DB)
	museumService = service.NewMuseumService(museumRepo)
}

func CreateGallery(c *gin.Context) {
	var req domain.CreateGalleryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, err.Error())
		return
	}

	gallery, err := museumService.CreateGallery(&req)
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, gallery.ToResponse())
}

func GetGallery(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	gallery, err := museumService.GetGallery(uint(id))
	if err != nil {
		response.Fail(c, http.StatusNotFound, "展厅不存在")
		return
	}

	response.Success(c, gallery.ToResponse())
}

func ListGalleries(c *gin.Context) {
	zone := domain.GalleryZone(c.Query("zone"))
	galleries, err := museumService.ListGalleries(zone)
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	result := make([]*domain.GalleryResponse, len(galleries))
	for i, g := range galleries {
		result[i] = g.ToResponse()
	}

	response.Success(c, gin.H{
		"list":  result,
		"total": len(result),
	})
}

func UpdateGallery(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req domain.CreateGalleryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, err.Error())
		return
	}

	gallery, err := museumService.UpdateGallery(uint(id), &req)
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, gallery.ToResponse())
}

func DeleteGallery(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := museumService.DeleteGallery(uint(id)); err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, nil)
}

func CreateExhibit(c *gin.Context) {
	var req domain.CreateExhibitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, err.Error())
		return
	}

	exhibit, err := museumService.CreateExhibit(&req)
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, exhibit.ToResponse())
}

func GetExhibit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	exhibit, err := museumService.GetExhibit(uint(id))
	if err != nil {
		response.Fail(c, http.StatusNotFound, "展品不存在")
		return
	}

	response.Success(c, exhibit.ToResponse())
}

func ListExhibits(c *gin.Context) {
	galleryID, _ := strconv.Atoi(c.Query("galleryId"))
	level := domain.ExhibitLevel(c.Query("level"))
	exhibits, err := museumService.ListExhibits(uint(galleryID), level)
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	result := make([]*domain.ExhibitResponse, len(exhibits))
	for i, e := range exhibits {
		result[i] = e.ToResponse()
	}

	response.Success(c, gin.H{
		"list":  result,
		"total": len(result),
	})
}

func UpdateExhibit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req domain.CreateExhibitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, err.Error())
		return
	}

	exhibit, err := museumService.UpdateExhibit(uint(id), &req)
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, exhibit.ToResponse())
}

func DeleteExhibit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := museumService.DeleteExhibit(uint(id)); err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, nil)
}

func GetEnvironmentHistory(c *gin.Context) {
	galleryID, _ := strconv.Atoi(c.Param("galleryId"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))
	records, err := museumService.GetEnvironmentHistory(uint(galleryID), limit)
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	result := make([]*domain.EnvironmentRecordResponse, len(records))
	for i, r := range records {
		result[i] = r.ToResponse()
	}

	response.Success(c, gin.H{
		"list":  result,
		"total": len(result),
	})
}

func GetMuseumCheckItems(c *gin.Context) {
	response.Success(c, gin.H{
		"checkItems": domain.GalleryCheckItems,
		"pointTypes": []string{
			string(domain.PointTypeGallery),
			string(domain.PointTypeExhibit),
			string(domain.PointTypePublic),
			string(domain.PointTypeEquipment),
			string(domain.PointTypeSecurity),
		},
		"exhibitLevels": []string{
			string(domain.ExhibitLevelNational),
			string(domain.ExhibitLevelFirst),
			string(domain.ExhibitLevelSecond),
			string(domain.ExhibitLevelThird),
			string(domain.ExhibitLevelGeneral),
		},
		"galleryZones": []string{
			string(domain.GalleryZoneExhibition),
			string(domain.GalleryZoneStorage),
			string(domain.GalleryZonePublic),
			string(domain.GalleryZoneOffice),
			string(domain.GalleryZoneEquipment),
		},
	})
}
