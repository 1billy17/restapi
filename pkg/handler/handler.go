package handler

import (
	"TODOapi/pkg/service"
	"github.com/gin-gonic/gin"
)

type Hanler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Hanler {
	return &Hanler{services: services}
}

func (h *Hanler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sing-up", h.SingUp)
		auth.POST("/sing-in", h.SingIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getItems)
				items.GET("/:items_id", h.getItemById)
				items.PUT("/:items_id", h.updateItem)
				items.DELETE("/:items_id", h.deleteItem)

			}
		}
	}

	return router
}
