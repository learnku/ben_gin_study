package controller

import (
	"ben_gin_study/app/dbserver"
	"ben_gin_study/app/model"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (this *UserController) Get(ctx *gin.Context) {
	this.RespSuccess(ctx, gin.H{
		"ip": ctx.ClientIP(),
	}, nil)
}

func (this *UserController) Create(ctx *gin.Context) {
	mod := model.User{}
	server := new(dbserver.User)
	if err := ctx.ShouldBind(&mod); err != nil{
		this.RespError(ctx, err.Error())
	} else {
		id, err := server.Insert(mod)
		if err != nil{
			this.RespError(ctx, err.Error())
			return
		}
		this.RespCreated(ctx, id, "创建成功")
	}
}
