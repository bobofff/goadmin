package binding

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ParseError 将 Gin 的绑定/验证错误转换为可读的中文提示。
func ParseError(err error, req interface{}, c *gin.Context) string {
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) && len(validationErrors) > 0 {
		fieldErr := validationErrors[0]
		fieldName := resolveFieldName(fieldErr, req)
		return fmt.Sprintf("参数 %s 不合法", fieldName)
	}

	var unmarshalTypeErr *json.UnmarshalTypeError
	if errors.As(err, &unmarshalTypeErr) {
		return fmt.Sprintf("参数 %s 类型错误", unmarshalTypeErr.Field)
	}

	var numErr *strconv.NumError
	if errors.As(err, &numErr) {
		if field := findNumericFieldError(c, req); field != "" {
			return fmt.Sprintf("参数 %s 不合法", field)
		}
		return fmt.Sprintf("数值参数格式错误: %s", numErr.Num)
	}

	return err.Error()
}

func resolveFieldName(fieldErr validator.FieldError, req interface{}) string {
	fieldName := fieldErr.Field()
	typ := reflect.TypeOf(req)
	if typ == nil {
		return fieldName
	}
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return fieldName
	}

	if structField, ok := typ.FieldByName(fieldErr.StructField()); ok {
		if tag := firstTagValue(structField, []string{"json", "form"}); tag != "" {
			fieldName = tag
		}
	}
	return fieldName
}

func firstTagValue(field reflect.StructField, keys []string) string {
	for _, key := range keys {
		tag := field.Tag.Get(key)
		if tag == "" {
			continue
		}
		parts := strings.Split(tag, ",")
		if len(parts) > 0 && parts[0] != "-" && parts[0] != "" {
			return parts[0]
		}
	}
	return ""
}

func findNumericFieldError(c *gin.Context, req interface{}) string {
	typ := reflect.TypeOf(req)
	if typ == nil {
		return ""
	}
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return ""
	}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if !isNumericKind(field.Type.Kind()) {
			continue
		}
		name := firstTagValue(field, []string{"form", "json"})
		if name == "" {
			name = field.Name
		}
		rawValues := getRequestValues(c, name)
		if len(rawValues) == 0 {
			continue
		}
		for _, val := range rawValues {
			if val == "" {
				continue
			}
			if err := parseNumericValue(val, field.Type.Kind(), field.Type.Bits()); err != nil {
				return name
			}
		}
	}
	return ""
}

func getRequestValues(c *gin.Context, key string) []string {
	if c == nil {
		return nil
	}
	values := c.QueryArray(key)
	if len(values) == 0 {
		values = c.PostFormArray(key)
	}
	return values
}

func isNumericKind(kind reflect.Kind) bool {
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Uintptr, reflect.Float32, reflect.Float64:
		return true
	case reflect.Bool:
		return true
	default:
		return false
	}
}

func parseNumericValue(value string, kind reflect.Kind, bitSize int) error {
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		_, err := strconv.ParseInt(value, 10, bitSize)
		return err
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		_, err := strconv.ParseUint(value, 10, bitSize)
		return err
	case reflect.Float32, reflect.Float64:
		_, err := strconv.ParseFloat(value, bitSize)
		return err
	case reflect.Bool:
		_, err := strconv.ParseBool(value)
		return err
	default:
		return nil
	}
}
