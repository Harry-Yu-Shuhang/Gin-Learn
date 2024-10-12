package controllers

import (
	"gin-learn/config"
	"gin-learn/models"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u UserController) PostRegister(ctx *gin.Context) {
	//接受用户名，密码和确认密码
	username := ctx.DefaultPostForm("user_name", "") //获取POST表单的username参数，如果没有则返回空字符串
	password := ctx.DefaultPostForm("password", "")
	confirm_password := ctx.DefaultPostForm("confirm_password", "")
	err_id := config.RegisterErr

	// if username == "" || password == "" || confirmPassword == "" {
	// 	ReturnError(ctx, 4001, "用户名或密码不能为空")
	// 	return
	// }
	if username == "" {
		ReturnError(ctx, err_id, "用户名不能为空")
		return
	} else if password == "" {
		ReturnError(ctx, err_id, "密码不能为空")
		return
	} else if confirm_password == "" {
		ReturnError(ctx, err_id, "确认密码不能为空")
		return
	}

	if password != confirm_password {
		ReturnError(ctx, err_id, "密码和确认密码不一致")
		return
	}
	user, _ := models.GetUserInfoByUserName(username) //忽略error，因为如果用户不存在，error会有错误，但是注册的时候我们期望用户不存在
	if user.ID != 0 {                                 //初始化为0,如果用户存在，则ID不为0
		ReturnError(ctx, err_id, "用户名已存在")
		return
	}
	HashedPassword, _ := EnPwdCode([]byte(password)) //加密密码
	_, err := models.PostRegister(username, HashedPassword)
	if err != nil {
		ReturnError(ctx, err_id, "密码加密错误,请联系管理员")
		return
	}
	ReturnSuccess(ctx, 0, "注册成功", true, 1)
}

type UserApi struct {
	ID       int    `json:"id"`
	Username string `json:"user_name"`
}

func (u UserController) PostLogin(ctx *gin.Context) {
	//接受用户名和密码
	username := ctx.DefaultPostForm("user_name", "")
	password := ctx.DefaultPostForm("password", "")
	err_id := config.LoginErr

	if username == "" {
		ReturnError(ctx, err_id, "用户名不能为空")
		return
	} else if password == "" {
		ReturnError(ctx, err_id, "密码不能为空")
		return
	}

	user, _ := models.GetUserInfoByUserName(username)
	if user.ID == 0 {
		ReturnError(ctx, err_id, "用户名不存在")
		return
	}
	HashedPassword, _ := EnPwdCode([]byte(password)) //加密密码
	if user.Password != HashedPassword {
		ReturnError(ctx, err_id, "密码错误")
		return
	}
	data := UserApi{ID: user.ID, Username: user.Username}
	session := sessions.Default(ctx)
	session.Set("login:"+strconv.Itoa(user.ID), user.ID) //前面的id用来登陆，后面的id是Session的值
	session.Save()                                       //Session存储在Redis中
	ReturnSuccess(ctx, 0, "登录成功", data, 1)
}

//这里加结构体是为了不冲突，比如user.go和order.go中都有一样的方法GetAll,不用结构体会冲突。
//不加结构体的写法
// func GetUserInfo(ctx *gin.Context) {
// 	ReturnSuccess(ctx, 0, "success", "user info", 1)
// }

// func (u UserController) GetUserSingle(ctx *gin.Context) { //(u UserController)说明是UserController的一种方法
// 	idStr := ctx.Param("id") //获取形如user/info/id这种形式的id值
// 	name := ctx.Param("name")

// 	id, _ := strconv.Atoi(idStr) //转换成int类型
// 	user, _ := models.GetUserSingle(id)
// 	ReturnSuccess(ctx, 0, name, user, 1)
// }

// func (u UserController) GetUserListTest(ctx *gin.Context) {
// 	users, err := models.GetUserListTest("id > ?", 3)
// 	if err != nil {
// 		ReturnError(ctx, 4004, "查询多个User失败,没有相关数据")
// 		return
// 	}
// 	ReturnSuccess(ctx, 0, "查询多个User成功", users, 1)
// }

// func (u UserController) AddUser(ctx *gin.Context) {
// 	username := ctx.DefaultPostForm("username", "") //key,value初始化，检查请求有没有username这个key，如果没有则返回空字符串
// 	id, err := models.AddUser(username)
// 	if err != nil {
// 		ReturnError(ctx, 4002, "保存失败")
// 		return
// 	}
// 	ReturnSuccess(ctx, 0, "保存成功", id, 1)
// }

// func (u UserController) UpdateUser(ctx *gin.Context) {
// 	username := ctx.DefaultPostForm("username", "") //key,value初始化，检查请求有没有username这个key，如果没有则返回空字符串
// 	idStr := ctx.DefaultPostForm("id", "")
// 	id, _ := strconv.Atoi(idStr)
// 	err := models.UpdateUser(id, username)
// 	if err != nil {
// 		ReturnError(ctx, 4002, "更新失败")
// 		return
// 	}
// 	ReturnSuccess(ctx, 0, "更新成功", true, 1)
// }

// func (u UserController) DeleteUser(ctx *gin.Context) {
// 	idStr := ctx.DefaultPostForm("id", "")
// 	id, _ := strconv.Atoi(idStr)
// 	err := models.DeleteUser(id)
// 	if err != nil {
// 		ReturnError(ctx, 4002, "删除失败")
// 		return
// 	}
// 	ReturnSuccess(ctx, 0, "删除成功", true, 1)
// }
