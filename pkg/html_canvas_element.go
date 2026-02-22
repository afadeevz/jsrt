//go:build js && wasm

package jsrt

import "syscall/js"

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLCanvasElement
type HTMLCanvasElement interface {
	HTMLElement

	// attributes
	GetWidth() int
	SetWidth(int)
	GetHeight() int
	SetHeight(int)

	// CaptureStream(frameRate ...float64) MediaStream
	GetContext2D(attributes ...CanvasRenderingContext2DAttributes) CanvasRenderingContext2D
	ToDataURL() string
	ToDataURL1(type_ string) string
	ToDataURL2(type_ string, quality float64) string
	// ToBlob(callback, type, quality) Blob
	// TransferControlToOffscreen() OffscreenCanvas
}

var _ HTMLCanvasElement = (*htmlCanvasElement)(nil)

func (c *htmlCanvasElement) GetContext2D(attributes ...CanvasRenderingContext2DAttributes) CanvasRenderingContext2D {
	if len(attributes) > 0 {
		attr := attributes[0]
		attrObj := js.ValueOf(map[string]any{
			"alpha":              attr.Alpha,
			"desynchronized":     attr.Desynchronized,
			"willReadFrequently": attr.WillReadFrequently,
		})
		if attr.colorSpace != nil {
			attrObj.Set("colorSpace", *attr.colorSpace)
		}
		if attr.colorType != nil {
			attrObj.Set("colorType", *attr.colorType)
		}
		return wrapCanvasRenderingContext2D(c.Call("getContext", "2d", attrObj))
	}

	return wrapCanvasRenderingContext2D(c.Call("getContext", "2d"))
}

func (c *htmlCanvasElement) ToDataURL() string {
	return c.Call("toDataURL").String()
}

func (c *htmlCanvasElement) ToDataURL1(type_ string) string {
	return c.Call("toDataURL", type_).String()
}

func (c *htmlCanvasElement) ToDataURL2(type_ string, quality float64) string {
	return c.Call("toDataURL", type_, quality).String()
}
