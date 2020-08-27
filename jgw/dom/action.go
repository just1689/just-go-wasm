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
	util.AddEventListener(element, event, func(result []js.Value) {
		notify()
	})
	return
}
