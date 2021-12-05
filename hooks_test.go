package gromer

import (
	"context"
	"testing"

	. "github.com/franela/goblin"
)

func TestHooks(t *testing.T) {
	g := Goblin(t)
	g.Describe("useState", func() {
		ctx := WithState(context.Background())
		stateCtx := getState(ctx)
		getValue, setValue := UseState(ctx, 12)

		g.It("should be initialized ", func() {
			g.Assert(1).Equal(stateCtx.index)
			g.Assert(1).Equal(len(stateCtx.datas))
			g.Assert(stateCtx.datas[0]).Equal(12)
		})

		g.It("should get value ", func() {
			g.Assert(getValue()).Equal(12)
		})

		g.It("should set value", func() {
			setValue(15)
			g.Assert(stateCtx.datas[0]).Equal(15)
			g.Assert(getValue()).Equal(15)
		})
	})
}
