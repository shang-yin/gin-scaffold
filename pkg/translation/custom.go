package translation

import (
	"regexp"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func custom(v *validator.Validate) (err error) {
	// 自定义验证方法
	if err = v.RegisterValidation("checkMobile", checkMobile); err != nil {
		return err
	}
	// 自定义验证器
	if err = v.RegisterTranslation("checkMobile", trans,
		registerTranslator("checkMobile", "{0}格式错误!"),
		translate,
	); err != nil {
		return err
	}
	return nil
}

// registerTranslator 为自定义字段添加翻译功能
func registerTranslator(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(trans ut.Translator) error {
		if err := trans.Add(tag, msg, false); err != nil {
			return err
		}
		return nil
	}
}

// translate 自定义字段的翻译方法
func translate(trans ut.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		panic(fe.(error).Error())
	}
	return msg
}

func checkMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	re := `^1[3456789]\d{9}$`
	r := regexp.MustCompile(re)
	return r.MatchString(mobile)
}
