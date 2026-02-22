//go:build js && wasm

package jsrt

import "syscall/js"

// https://developer.mozilla.org/en-US/docs/Web/API/Node
type Node interface {
	EventTarget

	AppendChild(child Node) Node

	GetParentElement() Element
	FirstChild() Node
}

type node struct {
	*eventTarget
}

func newNode(value js.Value) *node {
	if value.IsNull() {
		return nil
	}
	return &node{newEventTarget(value)}
}

func wrapNode(value js.Value) Node {
	switch {
	case instanceOf(value, "Element"):
		return wrapElement(value)
	default:
		return newNode(value)
	}
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
