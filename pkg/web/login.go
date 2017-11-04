package main

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
	"github.com/oskca/gopherjs-vue"
)

type Model struct {
	*js.Object // this is needed for bidirectional data bindings

	Username  string `js:"username"`
	Password  string `js:"password"`
}

func (m *Model) Submit() {
	vue.GetVM(m).Call(
		"$alert",
		fmt.Sprintf("Username: %s", m.Username),
		map[string]string{
			"confirmButtonText": "OK",
		},
	)
}

func main() {
	m := &Model{
		Object: js.Global.Get("Object").New(),
	}

	// field assignment is required in this way to make data passing works
	m.Username = ""
	m.Password = ""

	// create the VueJS viewModel using a struct pointer
	vue.Use(js.Global.Get("ELEMENT"))
	vue.New("#app", m)
}
