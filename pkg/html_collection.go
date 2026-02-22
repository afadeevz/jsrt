//go:build js && wasm

package jsrt

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLCollection
type HTMLCollection interface {
	GetLength() uint
	Item(index uint) Element
	NamedItem(key string) Element
}

func (hc *htmlCollection) GetLength() uint {
	return uint(hc.Get("length").Int())
}

func (hc *htmlCollection) Item(index uint) Element {
	element := hc.Call("item", index)
	return wrapElement(element)
}

func (hc *htmlCollection) NamedItem(key string) Element {
	element := hc.Call("namedItem", key)
	return wrapElement(element)
}
