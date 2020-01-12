package router

import (
	"e-shop/Controllers"
	"e-shop/Web/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())

	v1 := r.Group("/")
	{
		v1.POST("register", Controllers.CreateUser)
		v1.POST("/login", Controllers.Login)

		v1.POST("/createproduct", Controllers.CreateProduct)
		v1.GET("/getproducts", Controllers.GetProducts)
		v1.PUT("/updateproduct/:id", Controllers.UpdateProduct)
		v1.GET("/getproduct/:id", Controllers.GetProduct)

	}

	return r
}

