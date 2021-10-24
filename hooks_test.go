package wapp

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestHooks(t *testing.T) {
	g := Goblin(t)
	g.Describe("useState", func() {
		ctx := &Context{index: 0, datas: []interface{}{}}
		getValue, setValue := UseState(ctx, 12)

		g.It("should be initialized ", func() {
			g.Assert(1).Equal(ctx.index)
			g.Assert(1).Equal(len(ctx.datas))
			g.Assert(ctx.datas[0]).Equal(12)
		})

		g.It("should get value ", func() {
			g.Assert(getValue()).Equal(12)
		})

		g.It("should set value", func() {
			setValue(15)
			g.Assert(ctx.datas[0]).Equal(15)
			g.Assert(getValue()).Equal(15)
		})
	})
}
