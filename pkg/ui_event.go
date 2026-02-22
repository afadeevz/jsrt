//go:build js && wasm

package jsrt

// https://developer.mozilla.org/en-US/docs/Web/API/UIEvent
type UIEvent interface {
	Event
}
