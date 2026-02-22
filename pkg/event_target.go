//go:build js && wasm

package jsrt

import "syscall/js"

type EventHandlerFn func(event Event)

// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget
type EventTarget interface {
	Wrapper

	AddEventListener(eventType string, listener EventHandlerFn) // TODO: opts
}

func (et *eventTarget) AddEventListener(eventType string, listener EventHandlerFn) {
	listenerImpl := js.FuncOf(func(_ js.Value, args []js.Value) any {
		go listener(wrapEvent(args[0]))
		return js.Undefined()
	})

	et.Call("addEventListener", eventType, listenerImpl)
}
