package validate

import (
	"sync"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/zeromicro/go-zero/core/logx"
)

// 定义一个全局翻译器T
var (
	trans        ut.Translator
	uni          *ut.UniversalTranslator
	validate     *validator.Validate
	validateOnce sync.Once
	transOnce    sync.Once
)

func Validate() *validator.Validate {
	validateOnce.Do(func() {
		validate = validator.New()
	})

	return validate
}

// InitTrans 初始化翻译器
func Trans(locale string) ut.Translator {
	transOnce.Do(func() {
		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器
		// 第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		// uni := ut.New(zhT, zhT) 也是可以的
		uni := ut.New(enT, zhT, enT)
		// locale 通常取决于 http 请求头的 'Accept-Language'
		var ok bool
		// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			logx.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		var err error
		// 注册翻译器
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(Validate(), trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(Validate(), trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(Validate(), trans)
		}

		if err != nil {
			logx.Errorf("register translator failed, err: %+v", err)
		}
	})

	return trans
}
