//go:build js && wasm

package jsrt

import "syscall/js"

// https://developer.mozilla.org/en-US/docs/Web/API/Window
type Window interface {
	RequestAnimationFrame(func(timestampMS float64))
}

func GetWindow() Window {
	return newWindow(js.Global())
}

// https://developer.mozilla.org/en-US/docs/Web/API/Window/requestAnimationFrame
func (w *window) RequestAnimationFrame(callback func(timestampMS float64)) {
	var release func()

	jsCallback := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		defer release()
		if len(args) != 1 {
			panic("expected exactly one argument for requestAnimationFrame callback")
		}
		timestamp := args[0].Float()
		go callback(timestamp)
		return nil
	})
	release = jsCallback.Release

	w.Call("requestAnimationFrame", jsCallback)
}
