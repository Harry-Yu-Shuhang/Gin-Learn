package controllers

import "github.com/gin-gonic/gin"

type UserController struct{}

//这里加结构体是为了不冲突，比如user.go和order.go中都有一样的方法GetAll,不用结构体会冲突。
//不加结构体的写法
// func GetUserInfo(ctx *gin.Context) {
// 	ReturnSuccess(ctx, 0, "success", "user info", 1)
// }

func (u UserController) GetUserInfo(ctx *gin.Context) {
	ReturnSuccess(ctx, 0, "success", "user info", 1)
}

func (u UserController) GetList(ctx *gin.Context) {
	ReturnError(ctx, 4004, "没有相关信息:user list")
}
