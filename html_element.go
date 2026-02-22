//go:build js && wasm

package jsrt

import "syscall/js"

type HTMLElement interface {
	Element
}

type htmlElement struct {
	*element
}

func newHTMLElement(value js.Value) *htmlElement {
	return &htmlElement{newElement(value)}
}

func wrapHTMLElement(value js.Value) HTMLElement {
	switch {
	case instanceOf(value, "HTMLInputElement"):
		return newHTMLInputElement(value)
	default:
		return newHTMLElement(value)
	}
}
