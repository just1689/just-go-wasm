package wswasm

import (
	"syscall/js"
)

func NewWSClient(url string, h WSInterface) WSClient {
	result := WSClient{
		url:      url,
		handlers: h,
	}
	return result
}

type WSClient struct {
	url      string
	ws       js.Value
	handlers WSInterface
}

type WSInterface interface {
	HandleMessage(string)
	HandleOnOpen()
	HandleClose()
}

func (w *WSClient) Send(s string) {
	w.ws.Call("send", s)
}

func (w *WSClient) Connect() {
	ws := js.Global().Get("WebSocket").New(w.url)
	ws.Call("addEventListener", "open", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		w.handlers.HandleOnOpen()
		return nil
	}))
	ws.Call("addEventListener", "message", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) < 1 {
			return nil
		}
		w.handlers.HandleMessage(args[0].Get("data").String())
		return nil
	}))
	ws.Call("addEventListener", "close", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		w.handlers.HandleClose()
		return nil
	}))

}
