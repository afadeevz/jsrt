//go:build js && wasm

package jsrt

import "syscall/js"

// https://developer.mozilla.org/en-US/docs/Web/API/Document
type Document interface {
	Wrapper

	CreateElement(tagName string) Element // TODO: options param
	CreateTextNode(text string) Text
	GetElementByID(id string) Element
	GetElementsByClassName(names string) HTMLCollection
}

func GetDocument() Document {
	return newDocument(js.Global().Get("document"))
}

type document struct {
	*wrapper
}

func newDocument(value js.Value) *document {
	if value.IsNull() {
		return nil
	}
	return &document{newWrapper(value)}
}

func (d *document) CreateElement(tagName string) Element {
	return wrapElement(d.Call("createElement", tagName))
}

func (d *document) CreateTextNode(text string) Text {
	return newText(d.Call("createTextNode", text))
}

func (d *document) GetElementByID(id string) Element {
	return wrapElement(d.Call("getElementById", id))
}

func (d *document) GetElementsByClassName(names string) HTMLCollection {
	return newHTMLCollection(d.Call("getElementsByClassName", names))
}
