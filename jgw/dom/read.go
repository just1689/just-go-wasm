package dom

import (
	"errors"
	"github.com/just1689/just-go-wasm/jgw/util"
	"reflect"
	"syscall/js"
)

func NewDOMReadBinder() (result *DOMReadBinder, err error) {
	result = &DOMReadBinder{
		jsDoc: js.Global().Get("document"),
	}
	if !result.jsDoc.Truthy() {
		err = errors.New("could not get document")
		return
	}
	return
}

type DOMReadBinder struct {
	jsDoc js.Value
}

func (d *DOMReadBinder) ReadDOMToStruct(i interface{}) {
	t := reflect.TypeOf(i)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag
		dom := tag.Get("dom")
		if dom != "" {
			element, err := util.GetElementByIDWithDoc(dom, d.jsDoc)
			if err != nil {
				continue
			}
			domValue := element.Get("value").String()
			ps := reflect.ValueOf(&i)
			structField := ps.FieldByName(field.Name)
			if structField.IsValid() && structField.CanSet() && structField.Kind() == reflect.String {
				structField.SetString(domValue)
			}
		}

	}
}
