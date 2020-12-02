package app

import (
	"encoding/json"
	"sync"

	"github.com/pyros2097/wapp/js"
)

func init() {
	LocalStorage = newJSStorage("localStorage")
	SessionStorage = newJSStorage("sessionStorage")
}

type jsStorage struct {
	name  string
	mutex sync.RWMutex
}

func newJSStorage(name string) *jsStorage {
	return &jsStorage{name: name}
}

func (s *jsStorage) Set(k string, v interface{}) (err error) {
	defer func() {
		r := recover()
		if r != nil {
			panic("setting storage value failed storage-type: " + s.name + "key " + k)
		}
	}()

	s.mutex.Lock()
	defer s.mutex.Unlock()

	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	js.Window.Get(s.name).Call("setItem", k, btos(b))
	return nil
}

func (s *jsStorage) Get(k string, v interface{}) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	item := js.Window.Get(s.name).Call("getItem", k)
	if !item.Truthy() {
		return nil
	}

	return json.Unmarshal(stob(item.String()), v)
}

func (s *jsStorage) Del(k string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	js.Window.Get(s.name).Call("removeItem", k)
}

func (s *jsStorage) Clear() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	js.Window.Get(s.name).Call("clear")
}
