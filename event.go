//go:build js && wasm

package jsrt

import "syscall/js"

// https://developer.mozilla.org/en-US/docs/Web/API/Event
type Event interface {
	GetBubbles() bool
	GetCancelable() bool
	GetComposed() bool
	GetCurrentTarget() EventTarget
	GetDefaultPrevented() bool
	// GetEventPhase()
	GetIsTrusted() bool
	GetTarget() EventTarget
	// GetTimeStamp()
	GetType() string
}

type event struct {
	*wrapper
}

func newEvent(value js.Value) *event {
	return &event{newWrapper(value)}
}

func (e *event) GetBubbles() bool {
	return e.Get("bubbles").Bool()
}

func (e *event) GetCancelable() bool {
	return e.Get("cancellable").Bool()
}

func (e *event) GetComposed() bool {
	return e.Get("composed").Bool()
}

func (e *event) GetCurrentTarget() EventTarget {
	return wrapEventTarget(e.Get("eventTarget"))
}

func (e *event) GetDefaultPrevented() bool {
	return e.Get("defaultPrevented").Bool()
}

func (e *event) GetIsTrusted() bool {
	return e.Get("isTrusted").Bool()
}

func (e *event) GetTarget() EventTarget {
	return wrapEventTarget(e.Get("target"))
}

func (e *event) GetType() string {
	return e.Get("type").String()
}
