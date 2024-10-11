package Res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}

func Success(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, http.StatusOK, data, msg)
}

func SuccessMsg(ctx *gin.Context, msg string) {
	Success(ctx, nil, msg)
}

func Fail(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusBadRequest, -1, data, msg)
}

func FailMsg(ctx *gin.Context, msg string) {
	Fail(ctx, nil, msg)
}
