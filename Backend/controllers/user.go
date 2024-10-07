package controllers

import (
	"gin-learn/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

//这里加结构体是为了不冲突，比如user.go和order.go中都有一样的方法GetAll,不用结构体会冲突。
//不加结构体的写法
// func GetUserInfo(ctx *gin.Context) {
// 	ReturnSuccess(ctx, 0, "success", "user info", 1)
// }

func (u UserController) GetUserInfo(ctx *gin.Context) {
	idStr := ctx.Param("id") //获取形如user/info/id这种形式的id值
	name := ctx.Param("name")

	id, _ := strconv.Atoi(idStr) //转换成int类型
	user, _ := models.GetUserTest(id)
	ReturnSuccess(ctx, 0, name, user, 1)
}

func (u UserController) GetList(ctx *gin.Context) {
	// logger.Write("日志信息", "user") //日志文件名以user开头
	// defer func() {
	// 	if err := recover(); err != nil { //加上后前端不返回错误，程序正常运行，但是后台打印
	// 		fmt.Println("捕获异常:", err)
	// 	}
	// }() //匿名函数，立即执行。如果不加最后这个括号，不会立即执行。
	num1 := 1
	num2 := 0
	num3 := num1 / num2
	ReturnError(ctx, 4004, num3)
}
