package gsx

import (
	"context"

	"github.com/jinzhu/copier"
)

type HX struct {
	Boosted     bool
	CurrentUrl  string
	Prompt      string
	Target      string
	TriggerName string
	TriggerID   string
}

type Context struct {
	context.Context
	hx      *HX
	data    M
	meta    M
	links   map[string]link
	scripts map[string]bool
}

func NewContext(c context.Context, hx *HX) *Context {
	return &Context{
		Context: c,
		hx:      hx,
		data:    M{},
		meta:    M{},
		links:   map[string]link{},
		scripts: map[string]bool{},
	}
}

func (c *Context) HX(k string) *HX {
	return c.hx
}

func (c *Context) Get(k string) interface{} {
	return c.data[k]
}

func (c *Context) Set(k string, v interface{}) {
	c.data[k] = v
}

func (c *Context) Meta(meta M) {
	c.meta = meta
}

func (c *Context) AddMeta(k, v string) {
	c.meta[k] = v
}

func (c *Context) Link(rel, href, t, as string) {
	c.links[href] = link{rel, href, t, as}
}

func (c *Context) Script(src string, sdefer bool) {
	c.scripts[src] = sdefer
}

func (c *Context) Data(data M) {
	c.data = data
}

func (c *Context) Render(tpl string) []*Tag {
	name, _ := c.Get("funcName").(string)
	tags := parse(name, tpl)
	return populate(c, tags)
}

func (c *Context) Clone(name string) *Context {
	clone := NewContext(c.Context, c.hx)
	err := copier.Copy(clone, c)
	if err != nil {
		panic("Failed to copy")
	}
	c.Set("funcName", name)
	return clone
}
