package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/assignment2-08/handlers"
)

func StartServer(h *handlers.Handlers) *gin.Engine {
	router := gin.Default()

	router.GET("/orders", h.OrderHandler.GetAllOrders)
	router.GET("/order/:id", h.OrderHandler.GetOrder)
	router.POST("/order", h.OrderHandler.CreateOrder)
	router.PUT("/order/:id", h.OrderHandler.UpdateOrder)
	router.PATCH("/order/:id", h.OrderHandler.UpdateOrder)
	router.DELETE("/order/:id", h.OrderHandler.DeleteOrder)

	return router
}
