package router

import (
	"gin-learn/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine { //要大写，公有，别的地方要用
	r := gin.Default()
	user := r.Group("/user")
	{
		user.GET("/info", controllers.UserController{}.GetUserInfo)
		user.POST("/list", controllers.UserController{}.GetList)
		user.PUT("/add", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "User Add")
		})
		user.DELETE("/delete", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "User Delete")
		})
	}

	order := r.Group("/order")
	{
		order.POST("/list", controllers.OrderController{}.GetList)
	}
	return r
}
