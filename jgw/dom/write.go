package dom

import (
	"github.com/just1689/just-go-wasm/jgw/contracts"
	"github.com/just1689/just-go-wasm/jgw/util"
	"syscall/js"
)

func NewWriteBinding(elementID string) (result *WriteBinding, err error) {
	return NewWriteBindingByField(elementID, "value")
}

func NewWriteBindingByField(elementID, field string) (result *WriteBinding, err error) {
	result = &WriteBinding{
		elementID: elementID,
		field:     field,
	}
	result.element, err = util.GetElementByID(elementID)
	if err != nil {
		return
	}
	result.Update = func(s string) { updater(result, s) }
	return
}

type WriteBinding struct {
	elementID string
	field     string
	lastValue string
	element   js.Value
	Update    contracts.PassString
}

func updater(w *WriteBinding, s string) {
	w.lastValue = s
	w.element.Set(w.field, s)
}
