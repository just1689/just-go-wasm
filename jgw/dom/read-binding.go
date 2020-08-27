package dom

import (
	"errors"
	"github.com/just1689/just-go-wasm/jgw/util"
	"log"
	"reflect"
	"syscall/js"
)

func NewDOMBinder(item *interface{}) (result *DOMMarshaller, err error) {
	result = &DOMMarshaller{
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

func (d *DOMMarshaller) bind(item *interface{}) {
	t := reflect.TypeOf(*item)
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
		util.AddEventListener(el, "change", func(result []js.Value) {
			ps := reflect.ValueOf(&i)
			structField := ps.FieldByName(field.Name)
			if structField.IsValid() && structField.CanSet() && structField.Kind() == reflect.String {
				domValue := el.Get("value").String() //Possibly use other fields
				structField.SetString(domValue)
			} else {
				log.Println("could not act on field by ", field.Name)
			}
		})
	}
}
