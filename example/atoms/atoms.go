package atoms

import (
	app "github.com/pyros2097/wapp"
)

var CountAtom = app.NewAtom(0)

func IncCount() {
	CountAtom.Set(CountAtom.Get().(int) + 1)
}

func DecCount() {
	CountAtom.Set(CountAtom.Get().(int) - 1)
}
