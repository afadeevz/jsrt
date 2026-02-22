//go:build js && wasm

package jsrt

import "syscall/js"

type EventHandlerFn func(event Event)

// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget
type EventTarget interface {
	Wrapper

	AddEventListener(eventType string, listener EventHandlerFn) // TODO: opts
}

type eventTarget struct {
	*wrapper
}

func newEventTarget(value js.Value) *eventTarget {
	if value.IsNull() {
		return nil
	}
	return &eventTarget{newWrapper(value)}
}

func wrapEventTarget(value js.Value) EventTarget {
	switch {
	case instanceOf(value, "Node"):
		return wrapNode(value)
	default:
		return newEventTarget(value)
	}
}

func (et *eventTarget) AddEventListener(eventType string, listener EventHandlerFn) {
	listenerImpl := js.FuncOf(func(_ js.Value, args []js.Value) any {
		go listener(newEvent(args[0]))
		return js.Undefined()
	})

	et.Call("addEventListener", eventType, listenerImpl)
}
