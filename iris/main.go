package main

import (
	"github.com/kataras/iris"
	"devx/iferr"
)

func main() {
	app := iris.New()

	app.Handle("POST", "/", func(ctx iris.Context) {
		var data struct {
			Name string
			Age  uint
		}
		iferr.Exit(ctx.ReadJSON(&data), "xxx")
		process()
		var ret struct {
			Result string
		}
		_, err := ctx.JSON(&ret)
		if err != nil {

		}
	})

	app.Run(iris.Addr(":8080"))
}
