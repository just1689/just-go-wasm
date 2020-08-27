package dom

import (
	"github.com/just1689/just-go-wasm/jgw/util"
	"syscall/js"
)

func BindEvent(elementID, event string, notify func()) (err error) {
	element, err := util.GetElementByID(elementID)
	if err != nil {
		return
	}
	element.Call("addEventListener", event, js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		notify()
		return nil
	}))
	return
}
