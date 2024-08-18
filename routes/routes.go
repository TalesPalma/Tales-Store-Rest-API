package routes

import (
	"github.com/TalesPalma/gin-golang-rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controllers.Index)
	r.GET("/products", controllers.GetProducts)
	r.GET("/products/:id", controllers.GetProductById)
	r.POST("products", controllers.PostProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)
	r.PUT("/products/:id", controllers.PutProduct)
	r.NoRoute(controllers.NotFound)
	r.Static("/assets", "./assets")
	r.Run()
}
