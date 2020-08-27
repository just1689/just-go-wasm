package dom

type HasSetter interface {
	Setter(field, value string)
}
