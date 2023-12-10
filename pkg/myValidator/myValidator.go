package myValidator

import (
	"fmt"
	"github.com/go-playground/locales/zh_Hans_CN"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/zh"
	"mountain/pkg/errcode"
	"reflect"
)

// 数据验证
func MyValidate(data interface{}) (string, int) {
	validate := validator.New()
	uni := ut.New(zh_Hans_CN.New())
	trans, _ := uni.GetTranslator("zh_Hans_CN")

	err := zh.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println("err:", err)
	}
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})

	err = validate.Struct(data) //检查结构体中的字段是否符合之前定义的验证规则
	if err != nil {

		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(trans), errmessage.ERROR
		}
	}
	return "", errmessage.SUCCESS
}
