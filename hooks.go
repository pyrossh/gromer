package gromer

import (
	"context"
)

type StateContext struct {
	index int
	datas []interface{}
}

// Go 1.18 stuff
// func UseState[T any](ctx *Context , s T) (func() T, func(v T)) {
// 	if len(ctx.datas) <= ctx.index + 1 {
// 			ctx .datas = append(ctx.datas, s)
// 	}
// 	return func() T {
// 	  return ctx.datas[ctx.index].(T)
// 	}, func(v T) {
// 	  ctx.datas[ctx.index] = v
// 	}
// }

func getState(ctx context.Context) *StateContext {
	return ctx.Value("state").(*StateContext)
}

func WithState(ctx context.Context) context.Context {
	return context.WithValue(ctx, "state", &StateContext{})
}

func UseState(ctx context.Context, s interface{}) (func() interface{}, func(v interface{})) {
	stateContext := getState(ctx)
	localIndex := stateContext.index
	if len(stateContext.datas) <= stateContext.index+1 {
		stateContext.datas = append(stateContext.datas, s)
		stateContext.index += 1
	}
	return func() interface{} {
			return stateContext.datas[localIndex].(interface{})
		}, func(v interface{}) {
			stateContext.datas[localIndex] = v
		}
}
