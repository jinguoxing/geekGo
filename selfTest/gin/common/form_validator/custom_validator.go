package form_validator

import (
    "github.com/go-playground/validator/v10"
    "time"
)

// customFunc 自定义字段级别校验方法
func customFunc(fl validator.FieldLevel) bool {
    date, err := time.Parse("2006-01-02", fl.Field().String())
    if err != nil {
        return false
    }
    if date.Before(time.Now()) {
        return false
    }
    return true
}
