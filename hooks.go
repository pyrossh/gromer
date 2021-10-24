package wapp

type Context struct {
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

func UseState(c *Context, s interface{}) (func() interface{}, func(v interface{})) {
	localIndex := c.index
	if len(c.datas) <= c.index+1 {
		c.datas = append(c.datas, s)
		c.index += 1
	}
	return func() interface{} {
			return c.datas[localIndex].(interface{})
		}, func(v interface{}) {
			c.datas[localIndex] = v
		}
}
