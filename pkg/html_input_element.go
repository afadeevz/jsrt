//go:build js && wasm

package jsrt

type HTMLInputElement interface {
	HTMLElement

	GetValue() string
	SetValue(string)
}
