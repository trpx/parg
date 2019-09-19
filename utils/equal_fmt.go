package utils

import (
	"fmt"
	"reflect"
)

func EqualFmt(a interface{}, b interface{}) bool {
	return reflect.TypeOf(a) == reflect.TypeOf(b) &&
		fmt.Sprintf("%#v", a) == fmt.Sprintf("%#v", b)
}
