package main

import (
	c "context"
	"io/ioutil"

	"wapp-example/context"
)

func main() {
	db := context.InitDB()
	ctx := c.Background()
	tx, err := context.BeginTransaction(db, ctx)
	if err != nil {
		panic(err)
	}
	files, err := ioutil.ReadDir("./migrations")
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		data, err := ioutil.ReadFile("./migrations/" + f.Name())
		if err != nil {
			panic(err)
		}
		tx.MustExec(string(data))
	}
	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}
