package embed

import (
	"errors"
	"reflect"
)

var ErrorNotEmbedded = errors.New("not embedded")
var ErrorUnableToSet = errors.New("unable to set")

func Embed(dst, src interface{}) error {
	if reflect.TypeOf(dst).Kind() != reflect.Ptr || reflect.TypeOf(src).Kind() != reflect.Ptr {
		return ErrorUnableToSet
	}
	return embed(reflect.ValueOf(dst).Elem(), reflect.ValueOf(src).Elem())
}

func embed(dst, src reflect.Value) error {
	dstKind := dst.Type().Kind()
	srcKind := src.Type().Kind()

	if dstKind == reflect.Struct && srcKind == reflect.Struct {
		return embedStruct(dst, src)
	}
	if dstKind == reflect.Struct || srcKind == reflect.Struct {
		return ErrorNotEmbedded
	}

	return embedField(dst, src)
}

func embedStruct(dst, src reflect.Value) error {
	srcType := src.Type()
	dstType := dst.Type()

	for i := 0; i < srcType.NumField(); i++ {
		fieldName := srcType.Field(i).Name

		if _, ok := dstType.FieldByName(fieldName); !ok {
			return ErrorNotEmbedded
		}

		if err := embed(dst.FieldByName(fieldName), src.Field(i)); err != nil {
			return err
		}
	}

	return nil
}

func embedField(dst, src reflect.Value) error {
	srcType := src.Type()
	dstType := dst.Type()

	if !srcType.ConvertibleTo(dstType) {
		return ErrorNotEmbedded
	}

	if !dst.CanSet() {
		return ErrorUnableToSet
	}

	dst.Set(src.Convert(dstType))

	return nil
}
