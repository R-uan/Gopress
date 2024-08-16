package gopress

import (
	"fmt"
	"reflect"
	"strconv"
)

func ParseContentLength(value string) int64 {
	if value == "" {
		return 0
	}
	length, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		fmt.Println("Error parsing Content-Length:", err)
		return 0
	}
	return length
}

func ParseHeadersToPlainText(v *ResponseHeaders) (string) {
 value := reflect.ValueOf(v)
    if value.Kind() == reflect.Ptr {
        value = value.Elem()
    }
    typ := value.Type()

    var result string
    ignoreEmpty := func(val reflect.Value) bool {
        switch val.Kind() {
        case reflect.String:
            return val.String() == ""
        case reflect.Int:
            return val.Int() == 0
        case reflect.Map:
            return val.Len() == 0
        case reflect.Slice:
            return val.Len() == 0
        default:
            return false
        }
    }

    for i := 0; i < value.NumField(); i++ {
        field := typ.Field(i)
        fieldValue := value.Field(i)

        if ignoreEmpty(fieldValue) {
            continue
        }

        if fieldValue.Kind() == reflect.Map {
            result += fmt.Sprintf("%s:\n", field.Name)
            for _, key := range fieldValue.MapKeys() {
                result += fmt.Sprintf("  %s: %v\n", key, fieldValue.MapIndex(key))
            }
        } else {
            result += fmt.Sprintf("%s: %v\n", field.Name, fieldValue)
        }
    }

    return result
}