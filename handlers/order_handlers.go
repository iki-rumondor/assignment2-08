package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/assignment2-08/applications/requests"
	"github.com/iki-rumondor/assignment2-08/applications/responses"
	"github.com/iki-rumondor/assignment2-08/applications/services"
	"github.com/iki-rumondor/assignment2-08/domains"
	"gorm.io/gorm"
)

type OrderHandler struct {
	Services *services.OrderServices
}

func (h *OrderHandler) GetAllOrders(c *gin.Context) {
	orders, err := h.Services.GetAllOrders()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Failed to get all orders",
		})
		return
	}

	var ordersRes = []responses.OrderRes{}

	for _, order := range *orders {

		var items []responses.Item

		for _, val := range order.Items {
			items = append(items, responses.Item{
				Id:          val.Id,
				CreatedAt:   val.CreatedAt,
				UpdatedAt:   val.UpdatedAt,
				ItemCode:    val.ItemCode,
				Description: val.Description,
				Quantity:    val.Quantity,
				OrderId:     val.OrderId,
			})
		}

		orderRes := responses.OrderRes{
			Id:           order.Id,
			CustomerName: order.CustomerName,
			CreatedAt:    order.CreatedAt,
			UpdatedAt:    order.UpdatedAt,
			Items:        items,
		}

		ordersRes = append(ordersRes, orderRes)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": ordersRes,
	})
}

func (h *OrderHandler) GetOrder(c *gin.Context) {

	orderId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Failed to parse order id",
		})
		return
	}

	order, err := h.Services.GetOrder(&orderId)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": fmt.Sprintf("Order with id %d Not Found", orderId),
		})
		return
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Failed to get order",
		})
		return
	}

	var items []responses.Item

	for _, val := range order.Items {
		items = append(items, responses.Item{
			Id:          val.Id,
			CreatedAt:   val.CreatedAt,
			UpdatedAt:   val.UpdatedAt,
			ItemCode:    val.ItemCode,
			Description: val.Description,
			Quantity:    val.Quantity,
			OrderId:     val.OrderId,
		})
	}

	orderRes := responses.OrderRes{
		Id:           order.Id,
		CustomerName: order.CustomerName,
		CreatedAt:    order.CreatedAt,
		UpdatedAt:    order.UpdatedAt,
		Items:        items,
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": orderRes,
	})
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var body requests.CreateOrder

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Failed to parse request body",
		})
		return
	}

	var items []domains.Item

	for _, val := range body.Items {
		items = append(items, domains.Item{
			ItemCode:    val.ItemCode,
			Description: val.Description,
			Quantity:    val.Quantity,
		})
	}

	mod := domains.Order{
		OrderedAt:    body.OrderedAt,
		CustomerName: body.CustomerName,
		Items:        items,
	}

	err := h.Services.CreateOrder(&mod)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"message": "Order has been created successfully",
	})
}

func (h *OrderHandler) UpdateOrder(c *gin.Context) {

	var body requests.UpdateOrder

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Failed to parse request body",
		})
		return
	}

	orderId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Failed to parse order id",
		})
		return
	}

	var items []domains.Item

	for _, val := range body.Items {
		items = append(items, domains.Item{
			ItemCode:    val.ItemCode,
			Description: val.Description,
			Quantity:    val.Quantity,
		})
	}

	mod := domains.Order{
		Id:           orderId,
		OrderedAt:    body.OrderedAt,
		CustomerName: body.CustomerName,
		Items:        items,
	}

	res, err := h.Services.UpdateOrder(&mod)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": fmt.Sprintf("Order with id %d Not Found", orderId),
		})
		return
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	var updatedItems []responses.Item

	for _, val := range res.Items {
		updatedItems = append(updatedItems, responses.Item{
			Id:          val.Id,
			CreatedAt:   val.CreatedAt,
			UpdatedAt:   val.UpdatedAt,
			ItemCode:    val.ItemCode,
			Description: val.Description,
			Quantity:    val.Quantity,
			OrderId:     val.OrderId,
		})
	}

	order := responses.OrderRes{
		Id:           res.Id,
		CustomerName: res.CustomerName,
		CreatedAt:    res.CreatedAt,
		UpdatedAt:    res.UpdatedAt,
		Items:        updatedItems,
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": order,
	})
}

func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	orderId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Failed to parse order id",
		})
		return
	}

	mod := domains.Order{
		Id: orderId,
	}

	err = h.Services.DeleteOrder(&mod)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": fmt.Sprintf("Order with id %d Not Found", orderId),
		})
		return
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Failed to update order",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": fmt.Sprintf("Order with id %d and items with order id %d has been deleted successfully", orderId, orderId),
	})

}
