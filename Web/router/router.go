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
		v1.POST("/subscribe", Controllers.Subscribe)
		v1.POST("/createaddress", Controllers.CreateAddress)
		v1.GET("/user_address/:id", Controllers.GetUserAddress)
		v1.POST("/stkpush", Controllers.MpesaExpress)
		v1.POST("/message", Controllers.SendMessages)
	}

	return r
}
