package model

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
)

// 表单验证
// https://github.com/go-playground/validator
var (
	uni      *ut.UniversalTranslator
	Validate *validator.Validate
	ValidateTrans ut.Translator
)

func init()  {
	zh2 := zh.New()
	uni = ut.New(zh2, zh2)

	ValidateTrans, _ = uni.GetTranslator("zh")

	Validate = validator.New()
	if err := zh_translations.RegisterDefaultTranslations(Validate, ValidateTrans); err != nil{
		logrus.Error(err)
	}
}

// 验证表单结构体
func ValidateModelFn(mod interface{}) (b bool, errMsg validator.ValidationErrorsTranslations) {
	err := Validate.Struct(mod)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		errMsg = errs.Translate(ValidateTrans)
	} else {
		b = true
	}
	return
}

func ValidateModelFn2(mod interface{}) (b bool) {
	err := Validate.Struct(mod)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			// can translate each error one at a time.
			fmt.Printf("%#v\n", e)
			e.Type().Elem()
			fmt.Printf("%#v\n", e.Translate(ValidateTrans))

		}
	} else {
		b = true
	}
	return
}