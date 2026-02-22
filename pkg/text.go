//go:build js && wasm

package jsrt

// https://developer.mozilla.org/en-US/docs/Web/API/Text
type Text interface {
	CharacterData
}
