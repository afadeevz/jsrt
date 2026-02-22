//go:build js && wasm

package jsrt

// https://developer.mozilla.org/en-US/docs/Web/API/DOMTokenList
type TokenList interface {
	Wrapper

	GetLength() uint
	GetValues() string
	Add(tokens ...string)
	Contains(token string) bool
	// Entries()
	// ForEach()
	Item(index uint) *string
	// Keys()
	Remove(tokens ...string)
	Replace(oldToken string, newToken string) bool
	Supports(token string) bool
	// Toggle()
	// Values()
}

func (tl *tokenList) Add(tokens ...string) {
	tl.Call("add", toSliceOfAny(tokens)...)
}

func (tl *tokenList) Contains(token string) bool {
	return tl.Call("contains", token).Bool()
}

func (tl *tokenList) Item(index uint) *string {
	res := tl.Call("item, index")
	if res.IsNull() {
		return nil
	}

	str := res.String()
	return &str
}

func (tl *tokenList) Remove(tokens ...string) {
	tl.Call("remove", toSliceOfAny(tokens)...)
}

func (tl *tokenList) Replace(oldToken string, newToken string) bool {
	return tl.Call("replace", oldToken, newToken).Bool()
}

func (tl *tokenList) Supports(token string) bool {
	return tl.Call("supports", token).Bool()
}
