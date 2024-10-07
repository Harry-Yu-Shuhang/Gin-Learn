package router

import (
	"gin-learn/controllers"
	"gin-learn/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine { //要大写，公有，别的地方要用
	r := gin.Default()

	//调用写日志函数,logger是自己写的
	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)

	user := r.Group("/user")
	{
		user.GET("/info/:id/:name", controllers.UserController{}.GetUserInfo) //:id可以通过ctx传入id
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
