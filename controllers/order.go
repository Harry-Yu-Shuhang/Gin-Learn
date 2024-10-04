package controllers

import "github.com/gin-gonic/gin"

type OrderController struct{}

func (u OrderController) GetOrderInfo(ctx *gin.Context) {
	ReturnSuccess(ctx, 0, "success", "order info", 1)
}

func (u OrderController) GetList(ctx *gin.Context) {
	ReturnError(ctx, 4004, "没有相关信息:order list")
}
