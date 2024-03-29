package gsx

import (
	"context"
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
	styles  M
}

func NewContext(c context.Context, hx *HX) *Context {
	return &Context{
		Context: c,
		hx:      hx,
		data:    M{},
		meta:    M{},
		links:   map[string]link{},
		scripts: map[string]bool{},
		styles:  M{},
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

func (c *Context) Styles(s M) {
	c.styles = s
}

func (c *Context) Render(tpl string) []*Tag {
	name, ok := c.Get("funcName").(string)
	if !ok {
		panic("funcName is required")
	}
	tags := parse(name, tpl)
	return populate(c, tags)
}

func (c *Context) Clone(name string) *Context {
	newCtx := &Context{
		data: M{},
	}
	for k, v := range c.data {
		newCtx.data[k] = v
	}
	newCtx.Set("funcName", name)
	return newCtx
}
