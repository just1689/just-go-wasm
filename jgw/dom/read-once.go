package dom

import (
	"errors"
	"github.com/just1689/just-go-wasm/jgw/util"
	"reflect"
	"syscall/js"
)

func NewDOMMarshaller() (result *DOMMarshaller, err error) {
	result = &DOMMarshaller{
		jsDoc: js.Global().Get("document"),
	}
	if !result.jsDoc.Truthy() {
		err = errors.New("could not get document")
		return
	}
	return
}

type DOMMarshaller struct {
	jsDoc js.Value
}

func (d *DOMMarshaller) ReadAll(i interface{}, setter Setter) {
	t := reflect.TypeOf(i)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag
		dom := tag.Get(GoTagDOM)
		if dom != "" {
			element, err := util.GetElementByIDWithDoc(dom, d.jsDoc)
			if err != nil {
				continue
			}
			domValue := element.Get("value").String() //Possibly use other fields
			setter(field.Name, domValue)
		}
	}
}
