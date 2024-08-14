package routes

import (
	"github.com/TalesPalma/gin-golang-rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/", index)
	r.GET("/products", controllers.GetProducts)
	r.GET("/products/:id", controllers.GetProductById)
	r.POST("products", controllers.PostProduct)
	r.Run()
}

func index(ctx *gin.Context) {
	ctx.String(200, "Welcome to tales store API")
}
