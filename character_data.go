//go:build js && wasm

package jsrt

import "syscall/js"

// https://developer.mozilla.org/en-US/docs/Web/API/CharacterData
type CharacterData interface {
	Node
}

type characterData struct {
	*node
}

func newCharacterData(value js.Value) *characterData {
	if value.IsNull() {
		return nil
	}
	return &characterData{newNode(value)}
}
