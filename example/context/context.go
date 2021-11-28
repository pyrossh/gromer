package context

import (
	c "context"
)

func WithContext(ctx c.Context) (c.Context, error) {
	return c.WithValue(ctx, "userId", "123"), nil
}

func GetUserID(ctx c.Context) string {
	return ctx.Value("userId").(string)
}
