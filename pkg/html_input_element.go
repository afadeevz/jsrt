//go:build js && wasm

package jsrt

type HTMLInputElement interface {
	HTMLElement

	GetValue() string
	SetValue(string)
}

func (hie *htmlInputElement) GetValue() string {
	return hie.Get("value").String()
}

func (hie *htmlInputElement) SetValue(value string) {
	hie.Set("value", value)
}
