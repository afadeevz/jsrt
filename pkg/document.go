//go:build js && wasm

package jsrt

import "syscall/js"

// https://developer.mozilla.org/en-US/docs/Web/API/Document
type Document interface {
	Node

	CreateElement(tagName string) Element // TODO: options param
	CreateTextNode(text string) Text
	GetElementByID(id string) Element
	GetElementsByClassName(names string) HTMLCollection
}

func GetDocument() Document {
	return newDocument(js.Global().Get("document"))
}

func (d *document) CreateElement(tagName string) Element {
	return wrapElement(d.Call("createElement", tagName))
}

func (d *document) CreateTextNode(text string) Text {
	return wrapText(d.Call("createTextNode", text))
}

func (d *document) GetElementByID(id string) Element {
	return wrapElement(d.Call("getElementById", id))
}

func (d *document) GetElementsByClassName(names string) HTMLCollection {
	return wrapHTMLCollection(d.Call("getElementsByClassName", names))
}
