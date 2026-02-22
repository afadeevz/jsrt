//go:build js && wasm

package jsrt

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

func (e *element) Remove() {
	e.Call("remove")
}

func (e *element) GetClassList() TokenList {
	return newTokenList(e.Get("classList"))
}

func (e *element) SetInnerHTML(innerHTML string) {
	e.Set("innerHTML", innerHTML)
}

func (e *element) QuerySelector(selectors string) Element {
	return wrapElement(e.Call("querySelector", selectors))
}
