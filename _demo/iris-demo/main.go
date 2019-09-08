package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()
	htmlEngine := iris.HTML("./", ".html")
	app.RegisterView(htmlEngine)

	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("hello world")
	})
	app.Get("hello", func(ctx iris.Context) {
		ctx.ViewData("Title", "测试页面")
		ctx.ViewData("Content", "Hello")
		ctx.View("hello.html")
	})
	app.Run(iris.Addr(":8888"), iris.WithCharset("utf-8"))
}
