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
	// 绑定到结构体
	err := ctx.ShouldBind(&mod)
	if err != nil{
		this.RespError(ctx, err.Error())
		return
	}

	// 表单验证
	b, errMsg := model.ValidateModelFn(mod)
	if !b {
		this.RespError(ctx, errMsg)
	}
	/*err = model.Validate.Struct(mod)
	if err != nil {
		errs := err.(validator.ValidationErrors)

		this.RespError(ctx, errs.Translate(model.ValidateTrans))
		return
	}*/
	return

	// 插入到数据库
	id, err := server.Insert(mod)
	if err != nil{
		this.RespError(ctx, err.Error())
		return
	}
	this.RespCreated(ctx, id, "创建成功")
}
