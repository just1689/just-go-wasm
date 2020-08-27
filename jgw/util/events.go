package util

import "syscall/js"

func AddEventListener(element js.Value, event string, handler func(result []js.Value)) {
	element.Call("addEventListener", event, js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		handler(args)
		return nil
	}))
	return
}
