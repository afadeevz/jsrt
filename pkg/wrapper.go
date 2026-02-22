//go:build js && wasm

package jsrt

import "syscall/js"

type Wrapper interface {
	Unwrap() js.Value
}

type wrapper struct {
	js.Value
}

func newWrapper(value js.Value) *wrapper {
	if value.IsNull() {
		return nil
	}
	return &wrapper{value}
}

func (w *wrapper) Unwrap() js.Value {
	return w.Value
}
