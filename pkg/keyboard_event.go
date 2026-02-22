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

func (ke *keyboardEvent) GetLocation() KeyboardEventLocation {
	return KeyboardEventLocation(ke.Get("location").Int())
}
