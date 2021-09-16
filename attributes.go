package app

import (
	"strconv"
)

type Attribute struct {
	Key   string
	Value string
}

func ID(v string) Attribute {
	return Attribute{"id", v}
}

func Style(v string) Attribute {
	return Attribute{"style", v}
}

func Accept(v string) Attribute {
	return Attribute{"accept", v}
}

func AutoComplete(v bool) Attribute {
	return Attribute{"autocomplete", strconv.FormatBool(v)}
}

func Checked(v bool) Attribute {
	return Attribute{"checked", strconv.FormatBool(v)}
}

func Disabled(v bool) Attribute {
	return Attribute{"disabled", strconv.FormatBool(v)}
}

func Name(v string) Attribute {
	return Attribute{"name", v}
}

func Type(v string) Attribute {
	return Attribute{"type", v}
}

func Value(v string) Attribute {
	return Attribute{"value", v}
}

func Placeholder(v string) Attribute {
	return Attribute{"placeholder", v}
}

func Src(v string) Attribute {
	return Attribute{"src", v}
}

func Defer() Attribute {
	return Attribute{"defer", "true"}
}

func ViewBox(v string) Attribute {
	return Attribute{"viewBox", v}
}

func X(v string) Attribute {
	return Attribute{"x", v}
}

func Y(v string) Attribute {
	return Attribute{"y", v}
}

func Href(v string) Attribute {
	return Attribute{"href", v}
}

func Target(v string) Attribute {
	return Attribute{"target", v}
}

func Rel(v string) Attribute {
	return Attribute{"rel", v}
}

type CssAttribute struct {
	classes string
}

func Css(d string) CssAttribute {
	return CssAttribute{classes: d}
}

func CssIf(v bool, d string) CssAttribute {
	if v {
		return CssAttribute{classes: d}
	}
	return CssAttribute{}
}

func XData(v string) Attribute {
	return Attribute{"x-data", v}
}

func XText(v string) Attribute {
	return Attribute{"x-text", v}
}

func MergeAttributes(parent *Element, uis ...interface{}) *Element {
	elems := []*Element{}
	for _, v := range uis {
		switch c := v.(type) {
		case Attribute:
			parent.setAttr(c.Key, c.Value)
		case CssAttribute:
			if vv, ok := parent.attrs["classes"]; ok {
				parent.setAttr("class", vv+" "+c.classes)
			} else {
				parent.setAttr("class", c.classes)
			}
		case *Element:
			elems = append(elems, c)
		case nil:
			// dont need to add nil items
		default:
			// fmt.Printf("%v\n", v)
			panic("unknown type in render")
		}
	}
	if !parent.selfClosing {
		parent.body = elems
	}
	return parent
}
