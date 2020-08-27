package wswasm

import (
	"github.com/just1689/just-go-wasm/jgw/util"
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
	util.AddEventListener(ws, "open", func(result []js.Value) {
		w.handlers.HandleOnOpen()
	})
	util.AddEventListener(ws, "message", func(result []js.Value) {
		if len(result) < 1 {
			return
		}
		w.handlers.HandleMessage(result[0].Get("data").String())
	})
	util.AddEventListener(ws, "close", func(result []js.Value) {
		w.handlers.HandleClose()
	})

}
