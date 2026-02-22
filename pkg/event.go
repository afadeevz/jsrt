//go:build js && wasm

package jsrt

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

func (e *event) GetCurrentTarget() EventTarget {
	return wrapEventTarget(e.Get("currentTarget"))
}
func (e *event) GetTarget() EventTarget {
	return wrapEventTarget(e.Get("target"))
}
