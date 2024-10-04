package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonStruct struct {
	//json变量名用小写，私有
	Code  int   `json:"code"` //Code变量对应json的code字段
	Msg   any   `json:"msg"`
	Data  any   `json:"data"`
	Count int64 `json:"count"` //mysql取的数字
}

type JsonErrStruct struct {
	Code int `json:"code"`
	Msg  any `json:"msg"`
}

func ReturnSuccess(ctx *gin.Context, code int, msg any, data any, count int64) {
	json := &JsonStruct{Code: code, Msg: msg, Data: data, Count: count}
	ctx.JSON(http.StatusOK, json)
}

func ReturnError(ctx *gin.Context, code int, msg any) {
	json := &JsonErrStruct{Code: code, Msg: msg}
	ctx.JSON(http.StatusOK, json)
}
