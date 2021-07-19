package main

import (
	"net/http"

	"github.com/anangnovriadi/api-golang-product/controllers"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Test",
		})
	})
	r.GET("/product", controllers.ListProduct)
	r.POST("/product", controllers.CreateProduct)

	return r
}
