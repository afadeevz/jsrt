//go:build js && wasm

package jsrt

import "syscall/js"

// https://developer.mozilla.org/en-US/docs/Web/API/Text
type Text interface {
	CharacterData
}

type text struct {
	*characterData
}

func newText(value js.Value) *text {
	if value.IsNull() {
		return nil
	}
	return &text{newCharacterData(value)}
}
