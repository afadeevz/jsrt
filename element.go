//go:build js && wasm

package jsrt

import "syscall/js"

// https://developer.mozilla.org/en-US/docs/Web/API/Element
type Element interface {
	Node

	Remove()
	GetClassList() TokenList
	SetID(string)
	GetID() string
	SetInnerHTML(string)
	QuerySelector(string) Element
}

type element struct {
	*node
}

func newElement(value js.Value) *element {
	if value.IsNull() {
		return nil
	}
	return &element{newNode(value)}
}

func wrapElement(value js.Value) Element {
	switch {
	case instanceOf(value, "HTMLElement"):
		return wrapHTMLElement(value)
	default:
		return newElement(value)
	}
}

func (e *element) Remove() {
	e.Call("remove")
}

func (e *element) GetClassList() TokenList {
	return newTokenList(e.Get("classList"))
}

func (e *element) SetID(id string) {
	e.Set("id", id)
}

func (e *element) GetID() string {
	return e.Get("id").String()
}

func (e *element) SetInnerHTML(innerHTML string) {
	e.Set("innerHTML", innerHTML)
}

func (e *element) QuerySelector(selectors string) Element {
	return wrapElement(e.Call("querySelector", selectors))
}
