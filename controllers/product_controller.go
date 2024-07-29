package controllers

import (
	"app/config"
	"app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
}

func (con *ProductController) Index(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	order := c.DefaultQuery("order", "id")
	direction := c.DefaultQuery("direction", "asc")
	var products []models.Product
	config.DB.Order(order + " " + direction).
		Offset((page - 1) * size).
		Limit(size).
		Find(&products)
	c.JSON(http.StatusOK, products)
}
