package main

import (
	"github.com/warmans/vue"
)

const tmpl = `
<div>
  <p>{{ Message }}</p>
  <input v-model="Message">
<div>
`

type Data struct {
	Message string
}

func main() {
	vue.New(
		vue.El("#app"),
		vue.Template(tmpl),
		vue.Data(&Data{Message: "Hello WebAssembly!"}),
	)

	select {}
}
