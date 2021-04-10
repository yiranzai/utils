package translator

import (
	"fmt"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	chTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var transMap = make(map[string]*ut.Translator)

//NewTranslator 返回一个翻译器
func NewTranslator(local string) (*ut.Translator, error) {
	if len(local) == 0 {
		local = "zh"
	}
	translator, ok := transMap[local]
	if !ok || translator == nil {
		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			zhT := zh.New() //chinese
			enT := en.New() //english
			uni := ut.New(enT, zhT, enT)

			t, o := uni.GetTranslator(local)
			if !o {
				return nil, fmt.Errorf("uni.GetTranslator(%s) failed", local)
			}
			//register translate
			// 注册翻译器
			switch local {
			case "en", "en-US":
				_ = enTranslations.RegisterDefaultTranslations(v, t)
			case "zh", "zh-CN":
				_ = chTranslations.RegisterDefaultTranslations(v, t)
			default:
				_ = enTranslations.RegisterDefaultTranslations(v, t)
			}
			transMap[local] = &t
		}
	}
	return transMap[local], nil
}
