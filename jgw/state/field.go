package state

import "github.com/just1689/just-go-wasm/jgw/contracts"

func NewRenderedString(initial string, updater contracts.PassString) *RenderedString {
	result := &RenderedString{
		s:       initial,
		updater: updater,
	}
	return result
}

type RenderedString struct {
	s       string
	updater contracts.PassString
}

func (r *RenderedString) Set(s string) {
	r.s = s
	r.updater(s)
}

func (r *RenderedString) Get() string {
	return r.s
}
