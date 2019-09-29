package controller

import (
	"ben_gin_study/app/dbserver"
	"ben_gin_study/app/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
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

func (this *UserController) Create(ctx *gin.Context) {
	mod := model.User{}
	if err := ctx.ShouldBind(&mod); err != nil{
		log.Fatal(err)
	} else {
		new(dbserver.User).Insert(&mod)
	}
}
