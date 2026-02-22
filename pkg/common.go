//go:build js && wasm

package jsrt

import (
	"fmt"
	"reflect"
	"syscall/js"
)

func valueOfFallback(r any, val any) js.Value {
	rVal := reflect.ValueOf(val)
	rType := reflect.TypeOf(val)

	switch rVal.Kind() {
	case reflect.Struct:
		ret := js.ValueOf(map[string]any{})
		for i := 0; i < rType.NumField(); i++ {
			rField := rType.Field(i)
			if !rField.IsExported() {
				continue
			}
			k := rField.Name
			v := rVal.Field(i).Interface()
			ret.Set(k, ValueOf(v))
		}
		return ret
	case reflect.Pointer:
		return ValueOf(rVal.Elem().Interface())
	case reflect.Array, reflect.Slice:
		ret := js.ValueOf([]any{})
		for i := 0; i < rVal.Len(); i++ {
			v := rVal.Index(i)
			ret.Call("push", ValueOf(v))
		}
		return ret
	case reflect.String:
		return js.ValueOf(rVal.String())
	}

	fmt.Println(rType)

	panic(r)
}

func ValueOf(val any) (ret js.Value) {
	defer func() {
		r := recover()
		if r != nil {
			ret = valueOfFallback(r, val)
		}
	}()

	return js.ValueOf(val)
}

func toSliceOfAny[T any](s []T) []any {
	result := make([]any, len(s))
	for i, v := range s {
		result[i] = v
	}
	return result
}

func instanceOf(value js.Value, className string) bool {
	if value.IsNull() || value.IsUndefined() {
		return false
	}

	class := js.Global().Get(className)
	return value.InstanceOf(class)
}

func fromOptionalString(value js.Value) *string {
	if value.IsNull() || value.IsUndefined() {
		return nil
	}
	str := value.String()
	return &str
}
