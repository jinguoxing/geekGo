package member

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/tidwall/gjson"
	"time"
)

//func NameValid(
//	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
//	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
//) bool {
//	if s, ok := field.Interface().(string); ok {
//		if s == "admin" {
//			return false
//		}
//	}
//	return true
//}


var NameValid validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

var SkuValid validator.Func = func(fl validator.FieldLevel) bool {

	skuJosn,ok := fl.Field().Interface().(string)

	if ok {

		value := gjson.Get(skuJosn,"sku")

	//	fmt.Println(value)

		return false

	}
	return true

}
