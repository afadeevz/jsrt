//go:build js && wasm

package jsrt

// https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D

// interface covering most of the 2D drawing API.  Parameters are deliberately
// weakly typed to make the wrapper easy to call; more-precise signatures can
// be added as needed.
type CanvasRenderingContext2D interface {
	ClearRect(x, y, w, h float64)
	FillRect(x, y, w, h float64)
	StrokeRect(x, y, w, h float64)

	FillText(text string, x, y float64)
	FillText1(text string, x, y float64, maxWidth float64)
	StrokeText(text string, x, y float64)
	StrokeText1(text string, x, y float64, maxWidth float64)
	// MeasureText(text string) TextMetrics

	GetLineWidth() float64
	SetLineWidth(float64)
	GetLineCap() string
	SetLineCap(string)
	GetLineJoin() string
	SetLineJoin(string)
	GetMiterLimit() float64
	SetMiterLimit(float64)
	GetLineDash() []float64
	SetLineDash([]float64)
	GetLineDashOffset() float64
	SetLineDashOffset(float64)

	GetFont() string
	SetFont(string)
	GetTextAlign() string
	SetTextAlign(string)
	GetTextBaseline() string
	SetTextBaseline(string)
	GetDirection() string
	SetDirection(string)
	GetLetterSpacing() string
	SetLetterSpacing(string)
	GetFontKerning() string
	SetFontKerning(string)
	GetFontStretch() string
	SetFontStretch(string)
	GetFontVariantCaps() string
	SetFontVariantCaps(string)
	GetTextRendering() string
	SetTextRendering(string)
	GetWordSpacing() string
	SetWordSpacing(string)
	GetLang() string
	SetLang(string)

	GetFillStyle() string
	SetFillStyle(string)
	GetStrokeStyle() string
	SetStrokeStyle(string)

	// CreateConicGradient(x, y, startAngle float64) any
	// CreateLinearGradient(x0, y0, x1, y1 float64) any
	// CreateRadialGradient(x0, y0, r0, x1, y1, r1 float64) any
	// CreatePattern(image any, repetition string) any

	GetShadowBlur() float64
	SetShadowBlur(float64)
	GetShadowColor() string
	SetShadowColor(string)
	GetShadowOffsetX() float64
	SetShadowOffsetX(float64)
	GetShadowOffsetY() float64
	SetShadowOffsetY(float64)

	BeginPath()
	ClosePath()
	MoveTo(x, y float64)
	LineTo(x, y float64)
	BezierCurveTo(cp1x, cp1y, cp2x, cp2y, x, y float64)
	QuadraticCurveTo(cpx, cpy, x, y float64)
	Arc(x, y, r, startAngle, endAngle float64, anticlockwise ...bool)
	ArcTo(x1, y1, x2, y2, r float64)
	Ellipse(x, y, radiusX, radiusY, rotation, startAngle, endAngle float64, anticlockwise ...bool)
	Rect(x, y, w, h float64)
	RoundRect(x, y, w, h, r float64)

	Fill()
	Stroke()
	// DrawFocusIfNeeded(el any)
	// Clip()
	IsPointInPath(x, y float64) bool
	IsPointInStroke(x, y float64) bool

	// GetTransform() any
	Rotate(angle float64)
	Scale(x, y float64)
	Translate(x, y float64)
	Transform(a, b, c, d, e, f float64)
	SetTransform(a, b, c, d, e, f float64)
	ResetTransform()

	GetGlobalAlpha() float64
	SetGlobalAlpha(float64)
	GetGlobalCompositeOperation() string
	SetGlobalCompositeOperation(string)

	// DrawImage(img any, args ...any)

	// CreateImageData(width, height float64) any
	// GetImageData(sx, sy, sw, sh float64) any
	// PutImageData(img any, dx, dy float64, args ...any)

	GetImageSmoothingEnabled() bool
	SetImageSmoothingEnabled(bool)
	GetImageSmoothingQuality() string
	SetImageSmoothingQuality(string)

	Save()
	Restore()
	Canvas() HTMLCanvasElement
	GetContextAttributes() CanvasRenderingContext2DAttributes
	Reset()
	IsContextLost() bool

	GetFilter() string
	SetFilter(string)
}

func (ctx *canvasRenderingContext2D) ClearRect(x, y, w, h float64) {
	ctx.Call("clearRect", x, y, w, h)
}
func (ctx *canvasRenderingContext2D) FillRect(x, y, w, h float64) {
	ctx.Call("fillRect", x, y, w, h)
}
func (ctx *canvasRenderingContext2D) StrokeRect(x, y, w, h float64) {
	ctx.Call("strokeRect", x, y, w, h)
}

func (ctx *canvasRenderingContext2D) FillText(text string, x, y float64) {
	ctx.Call("fillText", text, x, y)
}
func (ctx *canvasRenderingContext2D) FillText1(text string, x, y float64, maxWidth float64) {
	ctx.Call("fillText", text, x, y, maxWidth)
}
func (ctx *canvasRenderingContext2D) StrokeText(text string, x, y float64) {
	ctx.Call("strokeText", text, x, y)
}
func (ctx *canvasRenderingContext2D) StrokeText1(text string, x, y float64, maxWidth float64) {
	ctx.Call("strokeText", text, x, y, maxWidth)
}

func (ctx *canvasRenderingContext2D) GetLineDash() []float64 {
	arr := ctx.Call("getLineDash")
	n := arr.Length()
	res := make([]float64, n)
	for i := 0; i < n; i++ {
		res[i] = arr.Index(i).Float()
	}
	return res
}
func (ctx *canvasRenderingContext2D) SetLineDash(d []float64) {
	ctx.Call("setLineDash", toSliceOfAny(d))
}

func (ctx *canvasRenderingContext2D) BeginPath() {
	ctx.Call("beginPath")
}
func (ctx *canvasRenderingContext2D) ClosePath() {
	ctx.Call("closePath")
}
func (ctx *canvasRenderingContext2D) MoveTo(x, y float64) {
	ctx.Call("moveTo", x, y)
}
func (ctx *canvasRenderingContext2D) LineTo(x, y float64) {
	ctx.Call("lineTo", x, y)
}
func (ctx *canvasRenderingContext2D) BezierCurveTo(cp1x, cp1y, cp2x, cp2y, x, y float64) {
	ctx.Call("bezierCurveTo", cp1x, cp1y, cp2x, cp2y, x, y)
}
func (ctx *canvasRenderingContext2D) QuadraticCurveTo(cpx, cpy, x, y float64) {
	ctx.Call("quadraticCurveTo", cpx, cpy, x, y)
}
func (ctx *canvasRenderingContext2D) Arc(x, y, r, a1, a2 float64, ac ...bool) {
	if len(ac) > 0 {
		ctx.Call("arc", x, y, r, a1, a2, ac[0])
	} else {
		ctx.Call("arc", x, y, r, a1, a2)
	}
}
func (ctx *canvasRenderingContext2D) ArcTo(x1, y1, x2, y2, r float64) {
	ctx.Call("arcTo", x1, y1, x2, y2, r)
}
func (ctx *canvasRenderingContext2D) Ellipse(x, y, radiusX, radiusY, rotation, startAngle, endAngle float64, ac ...bool) {
	if len(ac) > 0 {
		ctx.Call("ellipse", x, y, radiusX, radiusY, rotation, startAngle, endAngle, ac[0])
	} else {
		ctx.Call("ellipse", x, y, radiusX, radiusY, rotation, startAngle, endAngle)
	}
}
func (ctx *canvasRenderingContext2D) Rect(x, y, w, h float64) {
	ctx.Call("rect", x, y, w, h)
}
func (ctx *canvasRenderingContext2D) RoundRect(x, y, w, h, r float64) {
	ctx.Call("roundRect", x, y, w, h, r)
}

func (ctx *canvasRenderingContext2D) Fill() {
	ctx.Call("fill")
}
func (ctx *canvasRenderingContext2D) Stroke() {
	ctx.Call("stroke")
}
func (ctx *canvasRenderingContext2D) IsPointInPath(x, y float64) bool {
	return ctx.Call("isPointInPath", x, y).Bool()
}
func (ctx *canvasRenderingContext2D) IsPointInStroke(x, y float64) bool {
	return ctx.Call("isPointInStroke", x, y).Bool()
}

func (ctx *canvasRenderingContext2D) Rotate(angle float64) {
	ctx.Call("rotate", angle)
}
func (ctx *canvasRenderingContext2D) Scale(x, y float64) {
	ctx.Call("scale", x, y)
}
func (ctx *canvasRenderingContext2D) Translate(x, y float64) {
	ctx.Call("translate", x, y)
}
func (ctx *canvasRenderingContext2D) Transform(a, b, c, d, e, f float64) {
	ctx.Call("transform", a, b, c, d, e, f)
}
func (ctx *canvasRenderingContext2D) SetTransform(a, b, c, d, e, f float64) {
	ctx.Call("setTransform", a, b, c, d, e, f)
}
func (ctx *canvasRenderingContext2D) ResetTransform() {
	ctx.Call("resetTransform")
}

/*
	Save()
	Restore()
	Canvas() HTMLCanvasElement
	GetContextAttributes() CanvasRenderingContext2DAttributes
	Reset()
	IsContextLost() bool
*/

type CanvasRenderingContext2DAttributes struct {
	Alpha              bool
	colorSpace         *string
	colorType          *string
	Desynchronized     bool
	WillReadFrequently bool
}

func (ctx *canvasRenderingContext2D) Save() {
	ctx.Call("save")
}
func (ctx *canvasRenderingContext2D) Restore() {
	ctx.Call("restore")
}
func (ctx *canvasRenderingContext2D) Canvas() HTMLCanvasElement {
	return wrapHTMLCanvasElement(ctx.Get("canvas"))
}
func (ctx *canvasRenderingContext2D) GetContextAttributes() CanvasRenderingContext2DAttributes {
	attr := ctx.Call("getContextAttributes")
	return CanvasRenderingContext2DAttributes{
		Alpha:              attr.Get("alpha").Bool(),
		Desynchronized:     attr.Get("desynchronized").Bool(),
		WillReadFrequently: attr.Get("willReadFrequently").Bool(),
		colorSpace:         fromOptionalString(attr.Get("colorSpace")),
		colorType:          fromOptionalString(attr.Get("colorType")),
	}
}
func (ctx *canvasRenderingContext2D) Reset() {
	ctx.Call("reset")
}
func (ctx *canvasRenderingContext2D) IsContextLost() bool {
	return ctx.Call("isContextLost").Bool()
}
