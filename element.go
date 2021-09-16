package app

import (
	"io"
	"unsafe"
)

type Element struct {
	tag         string
	attrs       map[string]string
	body        []*Element
	selfClosing bool
	text        string
}

func NewElement(tag string, selfClosing bool, uis ...interface{}) *Element {
	return MergeAttributes(&Element{tag: tag, selfClosing: selfClosing}, uis...)
}

func (e *Element) updateAttrs(attrs map[string]string) {
	for k := range e.attrs {
		if _, exists := attrs[k]; !exists {
			e.delAttr(k)
		}
	}

	if e.attrs == nil && len(attrs) != 0 {
		e.attrs = make(map[string]string, len(attrs))
	}

	for k, v := range attrs {
		if curval, exists := e.attrs[k]; !exists || curval != v {
			e.attrs[k] = v
		}
	}
}

func (e *Element) setAttr(k string, v string) {
	if e.attrs == nil {
		e.attrs = make(map[string]string)
	}

	switch k {
	case "style", "allow":
		s := e.attrs[k] + v + ";"
		e.attrs[k] = s
		return

	case "class":
		s := e.attrs[k]
		if s != "" {
			s += " "
		}
		s += v
		e.attrs[k] = s
		return
	}
	if v == "false" {
		delete(e.attrs, k)
		return
	} else if v == "true" {
		e.attrs[k] = ""
	} else {
		e.attrs[k] = v
	}
}

func (e *Element) delAttr(k string) {
	delete(e.attrs, k)
}

func writeIndent(w io.Writer, indent int) {
	for i := 0; i < indent*4; i++ {
		w.Write(stob(" "))
	}
}

func ln() []byte {
	return stob("\n")
}

func btos(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func stob(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func (e *Element) Html(w io.Writer) {
	e.HtmlWithIndent(w, 0)
}

func (e *Element) HtmlWithIndent(w io.Writer, indent int) {
	writeIndent(w, indent)
	if e.tag == "html" {
		w.Write(stob("<!DOCTYPE html>\n"))
	}
	if e.tag == "text" {
		writeIndent(w, indent)
		w.Write(stob(e.text))
		return
	}
	w.Write(stob("<"))
	w.Write(stob(e.tag))

	for k, v := range e.attrs {
		w.Write(stob(" "))
		w.Write(stob(k))

		if v != "" {
			w.Write(stob(`="`))
			w.Write(stob(v))
			w.Write(stob(`"`))
		}
	}

	w.Write(stob(">"))

	if e.selfClosing {
		return
	}

	for _, c := range e.body {
		w.Write(ln())
		if c != nil {
			c.HtmlWithIndent(w, indent+1)
		}
	}

	if len(e.body) != 0 {
		w.Write(ln())
		writeIndent(w, indent)
	}

	w.Write(stob("</"))
	w.Write(stob(e.tag))
	w.Write(stob(">"))
}
