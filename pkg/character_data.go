//go:build js && wasm

package jsrt

// https://developer.mozilla.org/en-US/docs/Web/API/CharacterData
type CharacterData interface {
	Node
}
