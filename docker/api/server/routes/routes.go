package routes

import (
	"gin-atomy-platform/server/service/products"

	"github.com/gin-gonic/gin"
)

// Engine - engine
func Engine() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")

	v1.GET("/products", products.GetProducts())

	return r
}
