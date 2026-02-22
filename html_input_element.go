//go:build js && wasm

package jsrt

import "syscall/js"

type HTMLInputElement interface {
	HTMLElement

	GetValue() string
	SetValue(string)
}

type htmlInputElement struct {
	*htmlElement
}

func newHTMLInputElement(value js.Value) *htmlInputElement {
	return &htmlInputElement{newHTMLElement(value)}
}

func (hie *htmlInputElement) GetValue() string {
	return hie.Get("value").String()
}

func (hie *htmlInputElement) SetValue(value string) {
	hie.Set("value", value)
}
