package config

import (
	"os"
	"reflect"
	"regexp"
)

var reVar = regexp.MustCompile(`^\${(\w+)}$`)

func fromenv(v interface{}) {
	fromenvReflect(reflect.ValueOf(v).Elem()) // assumes pointer to struct
}

func fromenvReflect(rv reflect.Value) {
	for i := 0; i < rv.NumField(); i++ {
		fv := rv.Field(i)

		if fv.Kind() == reflect.Ptr {
			fv = fv.Elem()
		}

		if fv.Kind() == reflect.Struct {
			fromenvReflect(fv)

			continue
		}

		if fv.Kind() == reflect.String {
			match := reVar.FindStringSubmatch(fv.String())
			if len(match) > 1 {
				fv.SetString(os.Getenv(match[1]))
			}
		}
	}
}
