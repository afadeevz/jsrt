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

func (e *event) GetBubbles() bool {
	return e.Get("bubbles").Bool()
}

func (e *event) GetCancelable() bool {
	return e.Get("cancelable").Bool()
}

func (e *event) GetComposed() bool {
	return e.Get("composed").Bool()
}

func (e *event) GetCurrentTarget() EventTarget {
	return wrapEventTarget(e.Get("currentTarget"))
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
