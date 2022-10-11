package form_validator

import (
    "fmt"
    "geekGo/selfTest/gin/vaildatorx"
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/gin/binding"
    ut "github.com/go-playground/universal-translator"
    "github.com/go-playground/validator/v10"
)

var (
    Validator *vaildatorx.CustomValidator
    // 定义一个全局翻译器T
)

func SetupValidator() error {
    Validator = vaildatorx.NewCustomValidator()
    Validator.Engine()
    binding.Validator = Validator

    return nil
}

func RegisterValidator(c *gin.Context) {
    v := c.Value("trans")
    trans, _ := v.(ut.Translator)

    if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
        if err := v.RegisterValidation("checkDate", customFunc); err != nil {

            fmt.Println(err)
            return
        }

        if err := v.RegisterTranslation(
            "checkDate",
            trans,
            registerTranslator("checkDate", "{0}必须要晚于当前日期"),
            translate,
        ); err != nil {
            fmt.Println(err)
        }

    }
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
