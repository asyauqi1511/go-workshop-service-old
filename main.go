package main

import (
	"github.com/asyauqi1511/go-workshop-service/controller/product"
	"github.com/asyauqi1511/go-workshop-service/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/products", product.Index)
	r.GET("/api/product/:id", product.Show)
	r.POST("/api/product", product.Create)
	r.PUT("/api/product/:id", product.Update)
	r.DELETE("/api/product", product.Delete)

	r.Run()
}
