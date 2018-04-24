package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	// Load all templates from the "./views" folder
	// where extension is ".html" and parse them
	// using the standard `html/template` package.
	app.RegisterView(iris.HTML("./views", ".html"))

	// Method:    GET
	// Resource:  http://localhost:8080
	app.Get("/", func(ctx iris.Context) {
		// Bind: {{.message}} with "Hello world!"
		ctx.ViewData("message", "Hello Iris!")
		// Render template file: ./views/index.html
		ctx.View("index.html")
	})

	// Method:    GET
	// Resource:  http://localhost:8080/column/666
	//
	// Need to use a custom regexp instead?
	// Easy,
	// just mark the parameter's type to 'string'
	// which accepts anything and make use of
	// its `regexp` macro function, i.e:
	// app.Get("/column/{id:string regexp(^[0-9]+$)}")
	app.Get("/column/{id:long}", func(ctx iris.Context) {
		columnID, _ := ctx.Params().GetInt64("id")
		ctx.Writef("Column ID: %d", columnID)
	})

	// Start the server using a network address.
	app.Run(iris.Addr(":8080"))
}
