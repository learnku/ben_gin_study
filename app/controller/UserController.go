package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (this *UserController) Get(ctx *gin.Context) {
	fmt.Printf("%#v", ctx.Request.Header.Get("User-Agent"))
	this.RespSuccess(ctx, gin.H{
		"ip": ctx.ClientIP(),
	}, nil)
}
