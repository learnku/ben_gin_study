package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct {
}

// 成功响应 200
func (this BaseController) RespSuccess(ctx *gin.Context, data, msg interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   data,
		"msg":    msg,
	})
}

// 失败响应 200
func (this BaseController) RespError(ctx *gin.Context, msg interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"msg":    msg,
	})
}

// 201 创建成功响应
func (this BaseController) RespCreated(ctx *gin.Context, id int64, msg interface{}) {
	if msg == nil {
		msg = "创建成功"
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"id":  id,
		"msg": msg,
	})
}

// 204 删除响应
func (this BaseController) RespDelete(ctx *gin.Context, msg interface{}) {
	ctx.JSON(http.StatusNoContent, nil)
}

// 401 用户未提供身份验证凭据，或者没有通过身份验证。
func (this BaseController) RespUnauthorized(ctx *gin.Context, msg interface{}) {
	if msg == nil {
		msg = "Unauthorized"
	}
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"msg": msg,
	})
}

// 403 用户通过了身份验证，但是不具有访问资源所需的权限。
func (this BaseController) RespForbidden(ctx *gin.Context, msg interface{}) {
	if msg == nil {
		msg = "Forbidden"
	}
	ctx.JSON(http.StatusForbidden, gin.H{
		"msg": msg,
	})
}

// 404 所请求的资源不存在，或不可用。
func (this BaseController) RespNotFound(ctx *gin.Context, msg interface{}) {
	if msg == nil {
		msg = "NotFound"
	}
	ctx.JSON(http.StatusNotFound, gin.H{
		"msg": msg,
	})
}

// 429 客户端的请求次数超过限额。
func (this BaseController) RespTooManyRequests(ctx *gin.Context, msg interface{}) {
	if msg == nil {
		msg = "TooManyRequests"
	}
	ctx.JSON(http.StatusTooManyRequests, gin.H{
		"msg": msg,
	})
}

//  500 客户端请求有效，服务器处理时发生了意外。
func (this BaseController) RespInternalServerError(ctx *gin.Context, msg interface{}) {
	if msg == nil {
		msg = "InternalServerError"
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"msg": msg,
	})
}

//  503 服务器无法处理请求，一般用于网站维护状态。
func (this BaseController) RespServiceUnavailable(ctx *gin.Context, msg interface{}) {
	if msg == nil {
		msg = "ServiceUnavailable"
	}
	ctx.JSON(http.StatusServiceUnavailable, gin.H{
		"msg": msg,
	})
}
