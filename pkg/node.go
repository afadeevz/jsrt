//go:build js && wasm

package jsrt

// https://developer.mozilla.org/en-US/docs/Web/API/Node
type Node interface {
	EventTarget

	AppendChild(child Node) Node

	GetParentElement() Element
	FirstChild() Node
}

func (n *node) AppendChild(child Node) Node {
	return wrapNode(n.Call("appendChild", child.Unwrap()))
}

func (n *node) GetParentElement() Element {
	return wrapElement(n.Get("parentElement"))
}

func (n *node) FirstChild() Node {
	return wrapNode(n.Get("firstChild"))
}
