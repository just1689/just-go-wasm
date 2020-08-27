package util

import (
	"errors"
	"fmt"
	"syscall/js"
)

func GetElementByID(elementID string) (element js.Value, err error) {
	jsDoc := js.Global().Get("document")
	if !jsDoc.Truthy() {
		err = errors.New("could not get document")
		return
	}
	element = jsDoc.Call("getElementById", elementID)
	if !element.Truthy() {
		err = errors.New(fmt.Sprint("could not get element by id: ", elementID))
		return
	}
	return
}
