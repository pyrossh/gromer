package main

import (
	c "context"

	"github.com/bxcodec/faker/v3"

	"wapp-example/context"
	"wapp-example/pages/api/todos"
)

func main() {
	db := context.InitDB()
	ctx := c.Background()
	tx, err := context.BeginTransaction(db, ctx)
	if err != nil {
		panic(err)
	}
	reqContext := context.ReqContext{
		Tx:     tx,
		UserID: "123",
	}
	for i := 0; i < 20; i++ {
		ti := todos.TodoInput{}
		err := faker.FakeData(&ti)
		if err != nil {
			panic(err)
		}
		_, _, err = todos.POST(reqContext, ti)
		if err != nil {
			panic(err)
		}
	}
	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}
