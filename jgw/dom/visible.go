package dom

import (
	"errors"
	"fmt"
	"github.com/just1689/just-go-wasm/jgw/util"
	"syscall/js"
)

func NewVisibilityBinding(elementID string) (result *VisibilityBinding, err error) {
	result = &VisibilityBinding{}
	e, err := util.GetElementByID(elementID)
	if err != nil {
		return
	}
	result.element = e.Get("style")
	if !result.element.Truthy() {
		err = errors.New(fmt.Sprint("could not get .style of ", elementID))
		return
	}
	return
}

type VisibilityBinding struct {
	element js.Value
}

func (v *VisibilityBinding) Set(visible bool) {
	if visible {
		v.element.Set("display", "block")
	} else {
		v.element.Set("display", "none")
	}
}
