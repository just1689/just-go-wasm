package dom

import (
	"errors"
	"fmt"
	"github.com/just1689/just-go-wasm/jgw/util"
	"log"
	"reflect"
	"syscall/js"
)

func NewDOMBinder(item HasSetter) (result *DOMBinder, err error) {
	result = &DOMBinder{
		jsDoc: js.Global().Get("document"),
	}
	if !result.jsDoc.Truthy() {
		err = errors.New("could not get document")
		return
	}
	result.bind(item)
	return
}

type DOMBinder struct {
	jsDoc js.Value
}

func (d *DOMBinder) bind(item HasSetter) {
	t := reflect.TypeOf(item)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag
		dom := tag.Get(GoTagDOM)
		if dom == "" {
			continue
		}
		el, err := util.GetElementByID(dom)
		if err != nil {
			log.Println("could not bind to field by ID ", dom)
			continue
		}
		fmt.Println("")
		util.AddEventListener(el, "change", func(_ []js.Value) {
			domValue := el.Get("value").String() //Possibly use other fields
			item.Setter(field.Name, domValue)
		})
	}
}
