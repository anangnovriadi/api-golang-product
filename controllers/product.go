package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/anangnovriadi/api-golang-product/models"
	"github.com/anangnovriadi/api-golang-product/utils"
	"github.com/gin-gonic/gin"
)

type ProductBody struct {
	Name        string  `json:"name,omitempty" binding:"required"`
	Price       float64 `json:"price,omitempty" binding:"required"`
	Description string  `json:"description,omitempty" binding:"required"`
	Quantity    int     `json:"quantity,omitempty" binding:"required"`
}

func ListProduct(c *gin.Context) {
	var order_by = c.Query("order_by")
	var price = c.Query("price")
	var name = c.Query("name")

	var orderBy = ""
	var queryPrice = ""
	var queryName = ""

	if order_by == "ASC" {
		orderBy = "asc"
	} else if order_by == "DESC" {
		orderBy = "desc"
	}

	if price == "ASC" {
		queryPrice = "asc"
	} else if price == "DESC" {
		queryPrice = "desc"
	}

	if name == "ASC" {
		queryName = "asc"
	} else if name == "DESC" {
		queryName = "desc"
	}

	var product []models.Products
	models.DB.Order("id " + orderBy).Order("price " + queryPrice).Order("name " + queryName).Find(&product)

	res, _ := json.Marshal(product)
	str := string(res)
	utils.Client("/product", str)

	c.JSON(http.StatusOK, gin.H{
		"message": "Found",
		"data":    product,
	})
}

func CreateProduct(c *gin.Context) {
	var body ProductBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := models.Products{Name: body.Name, Price: body.Price, Description: body.Description, Quantity: body.Quantity}
	models.DB.Create(&product)

	res, _ := json.Marshal(product)
	str := string(res)
	utils.Client("/product", str)

	c.JSON(http.StatusOK, gin.H{
		"message": "Created",
		"data":    product,
	})

}
