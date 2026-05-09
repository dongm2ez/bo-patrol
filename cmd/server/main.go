package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/bo-patrol/internal/pkg/config"
	"github.com/bo-patrol/internal/pkg/db"
	"github.com/bo-patrol/internal/middleware"
	"github.com/bo-patrol/internal/handler"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	if err := db.Init(); err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}

	handler.Init()

	gin.SetMode(config.GlobalConfig.Server.Mode)
	r := gin.New()

	r.Use(middleware.Logger())
	r.Use(middleware.CORS())
	r.Use(gin.Recovery())

	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", handler.Register)
			auth.POST("/login", handler.Login)
		}

		museum := api.Group("/museum")
		{
			museum.GET("/check-items", handler.GetMuseumCheckItems)
			museum.GET("/galleries", handler.ListGalleries)
			museum.GET("/exhibits", handler.ListExhibits)
		}

		public := api.Group("/public")
		{
			public.GET("/department-types", handler.GetDepartmentTypes)
			public.GET("/equipment-categories", handler.GetEquipmentCategories)
			public.GET("/space-types", handler.GetSpaceTypes)
			public.GET("/element-types", handler.GetElementTypes)
		}

		api.Use(middleware.JWTAuth())
		{
			user := api.Group("/users")
			{
				user.GET("", handler.ListUsers)
				user.GET("/:id", handler.GetUser)
				user.PUT("/:id", handler.UpdateUser)
				user.DELETE("/:id", handler.DeleteUser)
			}

			depts := api.Group("/departments")
			{
				depts.GET("", handler.ListDepartments)
				depts.GET("/tree", handler.GetDepartmentTree)
				depts.POST("", handler.CreateDepartment)
				depts.GET("/:id", handler.GetDepartment)
				depts.PUT("/:id", handler.UpdateDepartment)
				depts.DELETE("/:id", handler.DeleteDepartment)
			}

			equipments := api.Group("/equipments")
			{
				equipments.GET("", handler.ListEquipments)
				equipments.POST("", handler.CreateEquipment)
				equipments.GET("/:id", handler.GetEquipment)
				equipments.PUT("/:id", handler.UpdateEquipment)
				equipments.DELETE("/:id", handler.DeleteEquipment)
				equipments.GET("/department/:departmentId", handler.GetEquipmentByDepartment)
				equipments.GET("/:id/maintenance", handler.GetMaintenanceRecords)
			}

			suppliers := api.Group("/suppliers")
			{
				suppliers.GET("", handler.ListSuppliers)
				suppliers.POST("", handler.CreateSupplier)
			}

			spaces := api.Group("/spaces")
			{
				spaces.GET("", handler.ListLocations)
				spaces.GET("/tree", handler.GetLocationTree)
				spaces.POST("", handler.CreateLocation)
				spaces.GET("/:id", handler.GetLocation)
				spaces.PUT("/:id", handler.UpdateLocation)
				spaces.DELETE("/:id", handler.DeleteLocation)
			}

			mapElements := api.Group("/map-elements")
			{
				mapElements.GET("/location/:locationId", handler.GetMapElementsByLocation)
				mapElements.POST("", handler.CreateMapElement)
				mapElements.GET("/:id", handler.GetMapElement)
				mapElements.PUT("/:id", handler.UpdateMapElement)
				mapElements.DELETE("/:id", handler.DeleteMapElement)
				mapElements.PUT("/batch", handler.BatchUpdateMapElements)
			}

			museumAuth := api.Group("/museum")
			{
				gallery := museumAuth.Group("/galleries")
			{
				gallery.POST("", handler.CreateGallery)
				gallery.GET("/:id", handler.GetGallery)
				gallery.PUT("/:id", handler.UpdateGallery)
				gallery.DELETE("/:id", handler.DeleteGallery)
			}
			environment := museumAuth.Group("/environment")
			{
				environment.GET("/gallery/:id", handler.GetEnvironmentHistory)
			}

				exhibit := museumAuth.Group("/exhibits")
				{
					exhibit.POST("", handler.CreateExhibit)
					exhibit.GET("/:id", handler.GetExhibit)
					exhibit.PUT("/:id", handler.UpdateExhibit)
					exhibit.DELETE("/:id", handler.DeleteExhibit)
				}
			}

			routes := api.Group("/routes")
			{
				routes.GET("", handler.ListRoutes)
				routes.POST("", handler.CreateRoute)
				routes.GET("/:id", handler.GetRoute)
				routes.PUT("/:id", handler.UpdateRoute)
				routes.DELETE("/:id", handler.DeleteRoute)
				routes.GET("/:id/points", handler.ListPoints)
				routes.POST("/:id/points", handler.CreatePoint)
			}

			tasks := api.Group("/tasks")
			{
				tasks.GET("", handler.ListTasks)
				tasks.POST("", handler.CreateTask)
				tasks.GET("/:id", handler.GetTask)
				tasks.PUT("/:id/start", handler.StartTask)
				tasks.PUT("/:id/complete", handler.CompleteTask)
			}

			records := api.Group("/records")
			{
				records.GET("", handler.ListRecords)
				records.POST("", handler.CreateRecord)
				records.GET("/:id", handler.GetRecord)
			}

			tickets := api.Group("/tickets")
			{
				tickets.GET("", handler.ListTickets)
				tickets.POST("", handler.CreateTicket)
				tickets.GET("/:id", handler.GetTicket)
				tickets.PUT("/:id/assign", handler.AssignTicket)
				tickets.PUT("/:id/process", handler.ProcessTicket)
				tickets.PUT("/:id/complete", handler.CompleteTicket)
			}

			stats := api.Group("/stats")
			{
				stats.GET("/overview", handler.GetOverviewStats)
				stats.GET("/completion", handler.GetCompletionStats)
				stats.GET("/abnormal", handler.GetAbnormalStats)
			}
		}
	}

	addr := fmt.Sprintf(":%d", config.GlobalConfig.Server.Port)
	log.Printf("服务启动中，监听地址: %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
