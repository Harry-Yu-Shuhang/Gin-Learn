package controllers

import "github.com/gin-gonic/gin"

type OrderController struct{}

type Search struct {
	Name string `json:"name"`
	Cid  int    `json:"cid"`
}

func (u OrderController) GetOrderInfo(ctx *gin.Context) {
	ReturnSuccess(ctx, 0, "success", "order info", 1)
}

func (u OrderController) GetList(ctx *gin.Context) {
	//cid := ctx.PostForm("cid")                //从 POST 请求的表单 数据中获取字段 cid 的值。
	//name := ctx.DefaultPostForm("name", "王五") //读取name，默认王五
	// ReturnError(ctx, 4004, "没有相关信息:order list")
	//ReturnSuccess(ctx, 0, cid, name, 2)

	//接收json方式1:map
	// param := make(map[string]any)
	// err := ctx.BindJSON(&param)
	// if err != nil {
	// 	ReturnError(ctx, 4001, gin.H{"err": err}) //gin.H 是一个快捷方式，用于创建 map[string]any 类型的数据结构。它通常用于返回 JSON 格式的响应。
	// }
	// ReturnSuccess(ctx, 0, param["name"], param["cid"], 1)

	//接受方式2:结构体
	search := &Search{}
	err := ctx.BindJSON(&search)
	if err != nil {
		ReturnError(ctx, 4001, gin.H{"err": err}) //gin.H 是一个快捷方式，用于创建 map[string]any 类型的数据结构。它通常用于返回 JSON 格式的响应。
	}
	ReturnSuccess(ctx, 0, search.Name, search.Cid, 1)

}
