//go:build js && wasm

package jsrt

import "syscall/js"

// https://developer.mozilla.org/en-US/docs/Web/API/console
type Console interface {
	Wrapper

	Clear()
	Debug(data ...any)
	Error(data ...any)
	Info(data ...any)
	Warn(data ...any)
}

func GetConsole() Console {
	return newConsole(js.Global().Get("console"))
}

type console struct {
	*wrapper
}

func newConsole(value js.Value) Console {
	if value.IsNull() {
		return nil
	}
	return &console{newWrapper(value)}
}

func (c *console) Clear() {
	c.Call("clear")
}

func (c *console) Debug(data ...any) {
	c.Call("debug", data...)
}

func (c *console) Error(data ...any) {
	c.Call("error", data...)
}

func (c *console) Info(data ...any) {
	c.Call("info", data...)
}

func (c *console) Warn(data ...any) {
	c.Call("warn", data...)
}
