//go:build js && wasm

package jsrt

// https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent
type KeyboardEvent interface {
	UIEvent

	GetAltKey() bool
	GetCode() string
	GetCtrlKey() bool
	GetIsComposing() bool
	GetKey() string
	GetLocation() KeyboardEventLocation
	GetMetaKey() bool
	GetRepeat() bool
	GetShiftKey() bool
}

type KeyboardEventLocation uint

const (
	KeyboardEventLocationStandard KeyboardEventLocation = 0x00
	KeyboardEventLocationLeft     KeyboardEventLocation = 0x01
	KeyboardEventLocationRight    KeyboardEventLocation = 0x02
	KeyboardEventLocationNumpad   KeyboardEventLocation = 0x03
)

func (ke *keyboardEvent) GetAltKey() bool {
	return ke.Get("altKey").Bool()
}

func (ke *keyboardEvent) GetCode() string {
	return ke.Get("code").String()
}

func (ke *keyboardEvent) GetCtrlKey() bool {
	return ke.Get("ctrlKey").Bool()
}

func (ke *keyboardEvent) GetIsComposing() bool {
	return ke.Get("isComposing").Bool()
}

func (ke *keyboardEvent) GetKey() string {
	return ke.Get("key").String()
}

func (ke *keyboardEvent) GetLocation() KeyboardEventLocation {
	return KeyboardEventLocation(ke.Get("location").Int())
}

func (ke *keyboardEvent) GetMetaKey() bool {
	return ke.Get("metaKey").Bool()
}

func (ke *keyboardEvent) GetRepeat() bool {
	return ke.Get("repeat").Bool()
}

func (ke *keyboardEvent) GetShiftKey() bool {
	return ke.Get("shiftKey").Bool()
}
